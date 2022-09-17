package api

import (
	"strings"
	"sync"
)

type Extended struct {
	*Client
}

func FromClient(c *Client) *Extended {
	return &Extended{Client: c}
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
