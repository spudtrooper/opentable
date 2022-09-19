package api

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/spudtrooper/goutil/check"
	goutilerrors "github.com/spudtrooper/goutil/errors"
	goutillog "github.com/spudtrooper/goutil/log"
	"github.com/spudtrooper/goutil/or"
)

type Extended struct {
	*Client
	cache *Cache
}

func FromClient(c *Client, cache *Cache) *Extended {
	return &Extended{
		Client: c,
		cache:  cache,
	}
}

type FindMenuItemResultItem struct {
	Restaurant
	Items []RestaurantDetailsMenuSectionItem
}

type FindMenuItemInfo struct {
	Results []FindMenuItemResultItem
}

func (r FindMenuItemInfo) ForDebugging() FindMenuItemInfoForDebugging {
	var res FindMenuItemInfoForDebugging
	for _, it := range r.Results {
		res.Results = append(res.Results, FindMenuItemResultItemForDebugging{
			RestaurantName: it.Restaurant.Name,
			Items:          it.Items,
		})
	}
	return res
}

type FindMenuItemResultItemForDebugging struct {
	RestaurantName string
	Items          []RestaurantDetailsMenuSectionItem
}

type FindMenuItemInfoForDebugging struct {
	Results []FindMenuItemResultItemForDebugging
}

type FindMatchingMenuItemsInfo struct {
	Items []RestaurantDetailsMenuSectionItem
}

func FindMatchingMenuItems(rd RestaurantDetails, term string) FindMatchingMenuItemsInfo {
	var items []RestaurantDetailsMenuSectionItem
	for _, menu := range rd.Menus {
		for _, sec := range menu.Sections {
			for _, item := range sec.Items {
				if strings.EqualFold(item.Title, term) {
					items = append(items, item)
				}
			}
		}
	}
	return FindMatchingMenuItemsInfo{items}
}

func findMatchingMenuItems(rd RestaurantDetails, term string) []RestaurantDetailsMenuSectionItem {
	var items []RestaurantDetailsMenuSectionItem
	for _, menu := range rd.Menus {
		for _, sec := range menu.Sections {
			for _, item := range sec.Items {
				if strings.EqualFold(item.Title, term) {
					items = append(items, item)
				}
			}
		}
	}
	return items
}

type AllMenuItemsInfo struct {
	Items []RestaurantDetailsMenuSectionItem
}

func AllMenuItems(rd RestaurantDetails) AllMenuItemsInfo {
	var items []RestaurantDetailsMenuSectionItem
	for _, menu := range rd.Menus {
		for _, sec := range menu.Sections {
			items = append(items, sec.Items...)
		}
	}
	return AllMenuItemsInfo{items}
}

func allMenuItems(rd RestaurantDetails) []RestaurantDetailsMenuSectionItem {
	var items []RestaurantDetailsMenuSectionItem
	for _, menu := range rd.Menus {
		for _, sec := range menu.Sections {
			items = append(items, sec.Items...)
		}
	}
	return items
}

//go:generate genopts --function FindMenuItem verbose
func (e *Extended) FindMenuItem(term string, optss ...FindMenuItemOption) (*FindMenuItemInfo, error) {
	opts := MakeFindMenuItemOptions(optss...)

	searchInfo, err := e.Search(term, SearchVerbose(opts.Verbose()))
	if err != nil {
		return nil, err
	}

	itemsCh := make(chan FindMenuItemResultItem)
	errsCh := make(chan error)
	go func() {
		var wg sync.WaitGroup
		for _, r := range searchInfo.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				rd, err := e.RestaurantDetails(r, RestaurantDetailsVerbose(opts.Verbose()))
				if err != nil {
					errsCh <- err
					return
				}
				if items := findMatchingMenuItems(rd.RestaurantDetails, term); len(items) > 0 {
					itemsCh <- FindMenuItemResultItem{
						Restaurant: r,
						Items:      items,
					}
				}
			}()
		}
		wg.Wait()
		close(itemsCh)
		close(errsCh)
	}()

	var res FindMenuItemInfo
	for item := range itemsCh {
		res.Results = append(res.Results, item)
	}
	return &res, nil
}

//go:generate genopts --function RawSearchAllByURI verbose startPage:int
func (e *Extended) RawSearchAllByURI(uri string, optss ...RawSearchAllByURIOption) (chan RawSearchByURIInfo, chan error) {
	opts := MakeRawSearchAllByURIOptions(optss...)
	log := goutillog.MakeLog("RawSearchAllByURI")

	startPage := or.Int(opts.StartPage(), 1)

	// Remove the page=X query param. Retain the ? so we don't screw up the URI
	baseURI := uri
	baseURI = strings.Replace(baseURI, "&page=1", "&_page=1", 1)
	baseURI = strings.Replace(baseURI, "?page=1", "?_page=1", 1)

	log.Printf("searching starting with %q", baseURI)

	outCh, errCh := make(chan RawSearchByURIInfo), make(chan error)

	go func() {
		for page := startPage; ; page++ {
			uri := fmt.Sprintf("%s&page=%d", baseURI, page)
			log.Printf("searching %s", uri)
			info, err := e.RawSearchByURI(uri, SearchByURIVerbose(opts.Verbose()))
			if err != nil {
				errCh <- err
				break
			}
			if len(info.Convert().Restaurants) == 0 {
				log.Println("done")
				break
			}
			outCh <- *info
		}
		close(outCh)
		close(errCh)
	}()

	return outCh, errCh
}

//go:generate genopts --function RawListAllByURI verbose "startPage:int" "threads:int" "sync:bool"
func (e *Extended) RawListAllByURI(uri string, optss ...RawListAllByURIOption) (chan RawListByURIInfo, chan error) {
	opts := MakeRawListAllByURIOptions(optss...)

	startPage := or.Int(opts.StartPage(), 1)
	threads := or.Int(opts.Threads(), 5)
	verbose := opts.Verbose()

	// Remove the page=X query param. Retain the ? so we don't screw up the URI
	baseURI := uri
	baseURI = strings.Replace(baseURI, "&page=1", "&_page=1", 1)
	baseURI = strings.Replace(baseURI, "?page=1", "?_page=1", 1)

	if opts.Sync() {
		return e.rawListAllByURISync(baseURI, verbose, startPage)
	}
	return e.rawListAllByURIAsync(baseURI, verbose, startPage, threads)
}

func (e *Extended) listByURI(baseURI string, verbose bool, page int, outCh chan<- RawListByURIInfo, errCh chan<- error, log goutillog.Logger) (shouldBreak bool) {
	uri := fmt.Sprintf("%s&page=%d", baseURI, page)
	log.Printf("searching %s", uri)
	info, err := e.RawListByURI(uri, ListByURIVerbose(verbose))
	if err != nil {
		errCh <- err
		shouldBreak = true
		return
	}
	if len(info.Convert().Restaurants) == 0 {
		log.Println("done")
		shouldBreak = true
		return
	}
	outCh <- *info
	return
}

func (e *Extended) rawListAllByURISync(baseURI string, verbose bool, startPage int) (chan RawListByURIInfo, chan error) {
	log := goutillog.MakeLog("rawListAllByURISync")

	log.Printf("searching starting with %q", baseURI)

	outCh, errCh := make(chan RawListByURIInfo), make(chan error)

	go func() {
		for page := startPage; ; page++ {
			if shouldBreak := e.listByURI(baseURI, verbose, page, outCh, errCh, log); shouldBreak {
				break
			}
		}
		close(outCh)
		close(errCh)
	}()

	return outCh, errCh
}

func (e *Extended) rawListAllByURIAsync(baseURI string, verbose bool, startPage, threads int) (chan RawListByURIInfo, chan error) {
	log := goutillog.MakeLog("rawListAllByURIAsync")

	log.Printf("searching starting with %q", baseURI)

	inputCh := make(chan int)
	go func() {
		for page := startPage; ; page++ {
			inputCh <- page
		}
	}()

	outCh, errCh := make(chan RawListByURIInfo), make(chan error)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for page := range inputCh {
					if shouldBreak := e.listByURI(baseURI, verbose, page, outCh, errCh, log); shouldBreak {
						break
					}

				}
			}()
		}
		wg.Wait()
		close(outCh)
		close(errCh)
	}()

	return outCh, errCh
}

//go:generate genopts --function SearchAndSave verbose
func (e *Extended) SearchAndSave(ctx context.Context, term string, optss ...SearchAndSaveOption) error {
	opts := MakeSearchAndSaveOptions(optss...)

	info, err := e.Search(term, SearchVerbose(opts.Verbose()))
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	errsCh := make(chan error)
	go func() {
		defer wg.Done()
		var wg sync.WaitGroup
		for _, r := range info.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				rest, err := e.RawRestaurantDetailsByURI(r.ProfileLink)
				if err != nil {
					errsCh <- err
					return
				}
				if err := e.cache.SaveRestaurant(ctx, r.ProfileLink, *rest, SaveRestaurantVerbose(opts.Verbose())); err != nil {
					errsCh <- err
					return
				}
			}()
		}
		wg.Wait()
		close(errsCh)
	}()
	wg.Wait()

	eb := goutilerrors.MakeErrorCollector()
	for err := range errsCh {
		eb.Add(err)
	}
	if !eb.Empty() {
		return eb.Build()
	}
	return nil

}

//go:generate genopts --function SearchByURIAndSave verbose
func (e *Extended) SearchByURIAndSave(ctx context.Context, uri string, optss ...SearchByURIAndSaveOption) error {
	opts := MakeSearchByURIAndSaveOptions(optss...)

	info, err := e.SearchByURI(uri, SearchByURIVerbose(opts.Verbose()))
	if err != nil {
		return err
	}
	log.Printf("have %d restaurants", len(info.Restaurants))
	var wg sync.WaitGroup
	wg.Add(1)
	errsCh := make(chan error)
	go func() {
		defer wg.Done()
		var wg sync.WaitGroup
		for _, r := range info.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				rest, err := e.RawRestaurantDetailsByURI(r.ProfileLink)
				if err != nil {
					errsCh <- err
					return
				}
				if err := e.cache.SaveRestaurant(ctx, r.ProfileLink, *rest, SaveRestaurantVerbose(opts.Verbose())); err != nil {
					errsCh <- err
					return
				}
			}()
		}
		wg.Wait()
		close(errsCh)
	}()
	wg.Wait()

	eb := goutilerrors.MakeErrorCollector()
	for err := range errsCh {
		eb.Add(err)
	}
	if !eb.Empty() {
		return eb.Build()
	}
	return nil
}

//go:generate genopts --function AddRestaurantsToSearchByURIs "threads:int" verbose
func (e *Extended) AddRestaurantsToSearchByURIs(ctx context.Context, uri string, optss ...AddRestaurantsToSearchByURIsOption) error {
	opts := MakeAddRestaurantsToSearchByURIsOptions(optss...)

	rawInfos, errs := e.RawListAllByURI(uri, RawListAllByURIThreads(opts.Threads()))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var wg sync.WaitGroup
		for rawInfo := range rawInfos {
			info := rawInfo.Convert()
			for _, r := range info.Restaurants {
				r := r
				wg.Add(1)
				go func() {
					defer wg.Done()
					if err := e.cache.SaveRestaurantToSearch(ctx, r.ProfileLink, SaveRestaurantVerbose(opts.Verbose())); err != nil {
						check.Err(err)
						return
					}
				}()
			}
		}
		wg.Wait()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range errs {
			log.Printf("error: %v", e)
		}
	}()
	wg.Wait()
	return nil
}

//go:generate genopts --function SearchEmptyRestaurants "threads:int" verbose "sleep:time.Duration"
func (e *Extended) SearchEmptyRestaurants(ctx context.Context, optss ...SearchEmptyRestaurantsOption) error {
	opts := MakeSearchEmptyRestaurantsOptions(optss...)

	threads := or.Int(opts.Threads(), 20)
	sleep := or.Duration(opts.Sleep(), 3*time.Second)

	rs, err := e.cache.RestaurantsToSearch(ctx)
	if err != nil {
		return err
	}
	log.Printf("starting with %d restaurants to search", len(rs))
	resCh := make(chan string, len(rs))
	go func() {
		for _, r := range rs {
			resCh <- r.URI
		}
		close(resCh)
	}()

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		i := i
		wg.Add(1)
		go func() {
			log := goutillog.MakeLog(fmt.Sprintf("#%d", i))
			defer wg.Done()
			for uri := range resCh {
				log.Printf("searching %s", uri)
				rest, err := e.RawRestaurantDetailsByURI(uri)
				if err != nil {
					log.Printf("RawRestaurantDetailsByURI error: %v", err)
					return
				}
				if err := e.cache.SaveRestaurant(ctx, uri, *rest, SaveRestaurantVerbose(opts.Verbose())); err != nil {
					log.Printf("SaveRestaurant error: %v", err)
					return
				}
				if err := e.cache.DeleteRestaurantToSearch(ctx, uri, DeleteRestaurantVerbose(opts.Verbose())); err != nil {
					log.Printf("SaveRestaurant error: %v", err)
					return
				}
				if sleep != 0 {
					time.Sleep(sleep)
				}
			}
		}()
	}
	wg.Wait()

	return nil
}
