package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spudtrooper/goutil/flags"
	goutiljson "github.com/spudtrooper/goutil/json"
	minimalcli "github.com/spudtrooper/minimalcli/app"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/ingest"
)

var (
	verbose       = flags.Bool("verbose", "global verbose")
	term          = flags.String("term", "global term")
	debugFailures = flags.Bool("debug_failures", "global debug_failres")
	failureJSON   = flags.String("failure_json", "file to test parsing")
	restID        = flags.String("rest_id", `restaurant id, e.g. "kumi-japanese-restaurant-and-bar-nyc-new-york"`)
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
		if err := cache.SaveRestaurant(ctx, *info); err != nil {
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

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
