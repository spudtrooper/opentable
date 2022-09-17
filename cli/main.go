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
)

var (
	verbose       = flags.Bool("verbose", "global verbose")
	term          = flags.String("term", "global term")
	debugFailures = flags.Bool("debug_failures", "global debug_failres")
	failureJSON   = flags.String("failure_json", "file to test parsing")
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

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

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

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
