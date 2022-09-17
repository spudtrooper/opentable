package api

import (
	"log"
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
	RestaurantName string
	Items          []RestaurantDetailsMenuSectionItem
}

type FindMenuItemInfo struct {
	Results []FindMenuItemResultItem
}

//go:generate genopts --function FindMenuItem verbose
func (e *Extended) FindMenuItem(term string, optss ...FindMenuItemOption) (*FindMenuItemInfo, error) {
	opts := MakeFindMenuItemOptions(optss...)

	searchInfo, err := e.Search(term, SearchVerbose(opts.Verbose()))
	if err != nil {
		return nil, err
	}

	items := make(chan FindMenuItemResultItem)
	errs := make(chan error)
	go func() {
		var wg sync.WaitGroup
		for _, r := range searchInfo.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				var fmiri FindMenuItemResultItem
				rd, err := e.RestaurantDetails(r, RestaurantDetailsVerbose(opts.Verbose()))
				if err != nil {
					errs <- err
					return
				}
				for _, menu := range rd.RestaurantDetails.Menus {
					for _, sec := range menu.Sections {
						for _, item := range sec.Items {
							if strings.EqualFold(item.Title, term) {
								log.Printf("found item: %+v", item)
								fmiri.Items = append(fmiri.Items, item)
							}
						}
					}
				}
				if len(fmiri.Items) > 0 {
					fmiri.RestaurantName = r.Name
					items <- fmiri
				}
			}()
		}
		wg.Wait()
		close(items)
		close(errs)
	}()

	var res FindMenuItemInfo
	for item := range items {
		res.Results = append(res.Results, item)
	}
	return &res, nil
}
