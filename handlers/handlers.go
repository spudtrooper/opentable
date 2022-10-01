// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/apihandler"
)

func CreateHandlers(client *api.Extended) []apihandler.Handler {
	return []apihandler.Handler{

		apihandler.NewHandler(
			"Search",
			func(ctx apihandler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				info, err := client.Search(term, api.SearchVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		// TODO: Doesn't work
		apihandler.NewHandler(
			"LocationPicker",
			func(ctx apihandler.EvalContext) (interface{}, error) {
				info, err := client.LocationPicker(api.LocationPickerVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		apihandler.NewHandler(
			"RestaurantsAvailability",
			func(ctx apihandler.EvalContext) (interface{}, error) {
				info, err := client.RestaurantsAvailability(api.RestaurantsAvailabilityVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),
	}
}
