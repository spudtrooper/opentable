// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"github.com/spudtrooper/goutil/parallel"
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

		apihandler.NewHandler(
			"SearchAll",
			func(ctx apihandler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				infos, errs := client.SearchAll(term,
					api.SearchAllVerbose(ctx.Bool("verbose")),
					api.SearchAllThreads(ctx.Int("threads")),
					api.SearchAllStartPage(ctx.Int("start_page")),
				)
				type rest struct {
					Name, ProfileLink string
				}
				type resp struct {
					Restaurants []rest
					Errors      []error
				}
				var res resp
				parallel.WaitFor(
					func() {
						for info := range infos {
							for _, r := range info.SearchInfo.Restaurants {
								res.Restaurants = append(res.Restaurants, rest{
									Name:        r.Name,
									ProfileLink: r.ProfileLink,
								})
							}
						}
					},
					func() {
						for e := range errs {
							res.Errors = append(res.Errors, e)
						}
					})

				return res, nil
			}),
	}
}
