package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/flags"
	goutiljson "github.com/spudtrooper/goutil/json"
	minimalcli "github.com/spudtrooper/minimalcli/app"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/ingest"
)

var (
	verbose       = flags.Bool("verbose", "global verbose")
	term          = flags.String("term", "global term")
	uri           = flags.String("uri", "global uri")
	debugFailures = flags.Bool("debug_failures", "global debug_failres")
	failureJSON   = flags.String("failure_json", "file to test parsing")
	restID        = flags.String("rest_id", `restaurant id, e.g. "kumi-japanese-restaurant-and-bar-nyc-new-york"`)
	startPage     = flags.Int("start_page", "global start page")
)

func requireStringFlag(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}

func mustFormatString(x interface{}) string {
	return goutiljson.MustColorMarshal(x)
}

func Main(ctx context.Context) error {
	app := minimalcli.Make()
	app.Init()

	core, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}
	client := api.FromClient(core)

	db, err := ingest.ConnectToDB(ctx)
	if err != nil {
		return err
	}
	cache := ingest.MakeDBCache(db)

	app.Register("TestFailedJSON", func(context.Context) error {
		requireStringFlag(failureJSON, "failure_json")
		var payload interface{}
		log.Printf("testing failed json %q", *failureJSON)
		b, err := ioutil.ReadFile(*failureJSON)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &payload); err != nil {
			return err
		}
		return nil
	})

	app.Register("LocationPicker", func(context.Context) error {
		info, err := client.LocationPicker(api.LocationPickerVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("LocationPicker: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("HeaderUserProfile", func(context.Context) error {
		info, err := client.HeaderUserProfile(api.HeaderUserProfileVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("HeaderUserProfile: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("RestaurantsAvailability", func(context.Context) error {
		info, err := client.RestaurantsAvailability(api.RestaurantsAvailabilityVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("RestaurantsAvailability: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("Search", func(context.Context) error {
		requireStringFlag(term, "term")
		info, err := client.Search(*term, api.SearchVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("Search: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("SearchByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		info, err := client.SearchByURI(*uri, api.SearchByURIVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("SearchByURI: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("RawListByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		info, err := client.RawListByURI(*uri, api.ListByURIVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("RawListByURI: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("ListByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		info, err := client.ListByURI(*uri, api.ListByURIVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("ListByURI: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("RestaurantDetails", func(context.Context) error {
		requireStringFlag(term, "term")
		searchInfo, err := client.Search(*term, api.SearchVerbose(*verbose))
		if err != nil {
			return err
		}
		info, err := client.RestaurantDetails(searchInfo.Restaurants[0], api.RestaurantDetailsVerbose(*verbose), api.RestaurantDetailsDebugFailures(*debugFailures))
		if err != nil {
			return err
		}
		fmt.Printf("RestaurantDetails: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("RestaurantDetailsFromID", func(context.Context) error {
		requireStringFlag(restID, "rest_id")
		info, err := client.RestaurantDetailsFromID(*restID, api.RestaurantDetailsVerbose(*verbose), api.RestaurantDetailsDebugFailures(*debugFailures))
		if err != nil {
			return err
		}
		fmt.Printf("RestaurantDetailsFromID: %s\n", mustFormatString(info))
		return nil
	})

	app.Register("SaveRawRestaurantDetailsFromID", func(context.Context) error {
		requireStringFlag(restID, "rest_id")
		info, err := client.RawRestaurantDetailsFromID(*restID, api.RestaurantDetailsVerbose(*verbose), api.RestaurantDetailsDebugFailures(*debugFailures))
		if err != nil {
			return err
		}
		fmt.Printf("SaveRawRestaurantDetailsFromID: %s\n", mustFormatString(info))
		if err := cache.SaveRestaurant(ctx, *info, ingest.SaveRestaurantVerbose(*verbose)); err != nil {
			return err
		}
		return nil
	})

	app.Register("SearchAndSave", func(context.Context) error {
		requireStringFlag(term, "term")
		info, err := client.Search(*term, api.SearchVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("SearchAndSave: %s\n", mustFormatString(info))
		var wg sync.WaitGroup
		for _, r := range info.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				rest, err := client.RawRestaurantDetailsByURI(r.ProfileLink)
				if err != nil {
					check.Err(err)
					return
				}
				if err := cache.SaveRestaurant(ctx, *rest, ingest.SaveRestaurantVerbose(*verbose)); err != nil {
					check.Err(err)
					return
				}
			}()
		}
		wg.Wait()
		return nil
	})

	app.Register("SearchyURIAndSave", func(context.Context) error {
		requireStringFlag(uri, "uri")
		info, err := client.SearchByURI(*uri, api.SearchByURIVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("SearchyURIAndSave: %s\n", mustFormatString(info))
		log.Printf("have %d restaurants", len(info.Restaurants))
		var wg sync.WaitGroup
		for _, r := range info.Restaurants {
			r := r
			wg.Add(1)
			go func() {
				defer wg.Done()
				rest, err := client.RawRestaurantDetailsByURI(r.ProfileLink)
				if err != nil {
					check.Err(err)
					return
				}
				if err := cache.SaveRestaurant(ctx, *rest, ingest.SaveRestaurantVerbose(*verbose)); err != nil {
					check.Err(err)
					return
				}
			}()
		}
		wg.Wait()
		return nil
	})

	app.Register("FindMatchingMenuItems", func(context.Context) error {
		requireStringFlag(term, "term")
		requireStringFlag(restID, "rest_id")
		info, err := client.RestaurantDetailsFromID(*restID, api.RestaurantDetailsVerbose(*verbose), api.RestaurantDetailsDebugFailures(*debugFailures))
		if err != nil {
			return err
		}
		items := api.FindMatchingMenuItems(info.RestaurantDetails, *term)
		fmt.Printf("FindMatchingMenuItems: %s\n", mustFormatString(items))
		return nil
	})

	app.Register("AllMenuItems", func(context.Context) error {
		requireStringFlag(restID, "rest_id")
		info, err := client.RestaurantDetailsFromID(*restID, api.RestaurantDetailsVerbose(*verbose), api.RestaurantDetailsDebugFailures(*debugFailures))
		if err != nil {
			return err
		}
		items := api.AllMenuItems(info.RestaurantDetails)
		fmt.Printf("AllMenuItems: %s\n", mustFormatString(items))
		return nil
	})

	app.Register("FindMenuItem", func(context.Context) error {
		requireStringFlag(term, "term")
		info, err := client.FindMenuItem(*term, api.FindMenuItemVerbose(*verbose))
		if err != nil {
			return err
		}
		fmt.Printf("FindMenuItem: %s\n", mustFormatString(info.ForDebugging()))
		return nil
	})

	app.Register("RawSearchAllByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		rawInfos, errs := client.RawSearchAllByURI(*uri)
		if err != nil {
			return err
		}
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
						rest, err := client.RawRestaurantDetailsByURI(r.ProfileLink)
						if err != nil {
							check.Err(err)
							return
						}
						if err := cache.SaveRestaurant(ctx, *rest, ingest.SaveRestaurantVerbose(*verbose)); err != nil {
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
	})

	app.Register("RawListAllByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		rawInfos, errs := client.RawListAllByURI(*uri, api.RawListAllByURIStartPage(*startPage))
		if err != nil {
			return err
		}
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
						rest, err := client.RawRestaurantDetailsByURI(r.ProfileLink)
						if err != nil {
							check.Err(err)
							return
						}
						if err := cache.SaveRestaurant(ctx, *rest, ingest.SaveRestaurantVerbose(*verbose)); err != nil {
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
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
