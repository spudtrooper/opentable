package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spudtrooper/goutil/flags"
	goutiljson "github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/parallel"
	"github.com/spudtrooper/goutil/slice"
	minimalcli "github.com/spudtrooper/minimalcli/app"
	"github.com/spudtrooper/opentable/api"
)

var (
	verbose       = flags.Bool("verbose", "global verbose")
	term          = flags.String("term", "global term")
	uri           = flags.String("uri", "global uri")
	debugFailures = flags.Bool("debug_failures", "global debug_failres")
	failureJSON   = flags.String("failure_json", "file to test parsing")
	restID        = flags.String("rest_id", `restaurant id, e.g. "kumi-japanese-restaurant-and-bar-nyc-new-york"`)
	startPage     = flags.Int("start_page", "global start page")
	threads       = flags.Int("threads", "global threads")
	sleep         = flags.Duration("sleep", "global sleep")
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

	db, err := api.ConnectToDB(ctx)
	if err != nil {
		return err
	}
	cache := api.MakeDBCache(db)

	client := api.FromClient(core, cache)

	app.AddPostRunHook(func(ctx context.Context) error {
		log.Printf("\n\nStats\n%s", client.StatsString())
		return nil
	})

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
		uri := fmt.Sprintf("https://www.opentable.com/r/%s", *restID)
		if err := cache.SaveRestaurant(ctx, uri, *info, api.SaveRestaurantVerbose(*verbose)); err != nil {
			return err
		}
		return nil
	})

	app.Register("SearchAndSave", func(context.Context) error {
		requireStringFlag(term, "term")
		if err := client.SearchAndSave(ctx, *term, api.SearchAndSaveVerbose(*verbose)); err != nil {
			return err
		}
		return nil
	})

	app.Register("SearchByURIAndSave", func(context.Context) error {
		requireStringFlag(uri, "uri")
		if err := client.SearchByURIAndSave(ctx, *uri, api.SearchByURIAndSaveVerbose(*verbose)); err != nil {
			return err
		}
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

	app.Register("AddRestaurantsToSearchByURIs", func(context.Context) error {
		requireStringFlag(uri, "uri")
		uris := slice.NonEmptyStrings(slice.Strings(*uri, ","))
		for _, uri := range uris {
			toSearch, errs, err := client.AddRestaurantsToSearchByURIs(ctx, uri,
				api.AddRestaurantsToSearchByURIsThreads(*threads),
				api.AddRestaurantsToSearchByURIsVerbose(*verbose))
			if err != nil {
				return err
			}
			parallel.WaitFor(
				func() {
					for r := range toSearch {
						if err := client.SearchRestaurantFromQueue(ctx, r,
							api.SearchRestaurantFromQueueVerbose(*verbose),
						); err != nil {
							fmt.Printf("SearchRestaurantFromQueue: %v\n", err)
						}
					}
				},
				func() {
					for e := range errs {
						fmt.Printf("AddRestaurantsToSearchByURIs error: %v\n", e)
					}
				})
		}
		return nil
	})

	app.Register("SearchEmptyRestaurants", func(context.Context) error {
		if err := client.SearchEmptyRestaurants(ctx,
			api.SearchEmptyRestaurantsThreads(*threads),
			api.SearchEmptyRestaurantsVerbose(*verbose),
			api.SearchEmptyRestaurantsSleep(*sleep)); err != nil {
			return err
		}
		return nil
	})

	app.Register("RawListAllByURI", func(context.Context) error {
		requireStringFlag(uri, "uri")
		res, err := client.RawListAllByURI(*uri, api.RawListAllByURIVerbose(*verbose))
		parallel.WaitFor(
			func() {
				for r := range res {
					fmt.Printf("RawListAllByURI: %s\n", mustFormatString(r))
				}
			},
			func() {
				for e := range err {
					fmt.Printf("RawListAllByURI error: %v\n", e)
				}
			})
		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
