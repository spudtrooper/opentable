// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"fmt"
	"sync"
	"time"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/parallel"
	"github.com/spudtrooper/goutil/slice"
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opentable/api"
)

func CreateHandlers(client *api.Extended) []handler.Handler {
	return []handler.Handler{

		handler.NewHandler("Search",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				info, err := client.Search(term, api.SearchVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}, handler.NewHandlerMetadata(handler.HandlerMetadata{
				Params: []handler.HandlerMetadataParam{
					{
						Name: "term",
						Type: handler.HandlerMetadataParamTypeString,
					},
				},
			})),

		// TODO: Doesn't work
		handler.NewHandler("LocationPicker",
			func(ctx handler.EvalContext) (interface{}, error) {
				info, err := client.LocationPicker(api.LocationPickerVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("RestaurantsAvailability",
			func(ctx handler.EvalContext) (interface{}, error) {
				info, err := client.RestaurantsAvailability(api.RestaurantsAvailabilityVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("SearchAll",
			func(ctx handler.EvalContext) (interface{}, error) {
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

		handler.NewHandler("SearchByURI",
			func(ctx handler.EvalContext) (interface{}, error) {
				uri, ok := ctx.MustString("uri")
				if !ok {
					return nil, nil
				}
				info, err := client.SearchByURI(uri, api.SearchByURIVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("RawListByURI",
			func(ctx handler.EvalContext) (interface{}, error) {
				uri, ok := ctx.MustString("uri")
				if !ok {
					return nil, nil
				}
				info, err := client.RawListByURI(uri, api.ListByURIVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("ListByURI",
			func(ctx handler.EvalContext) (interface{}, error) {
				uri, ok := ctx.MustString("uri")
				if !ok {
					return nil, nil
				}
				info, err := client.ListByURI(uri, api.ListByURIVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("RestaurantDetails",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				searchInfo, err := client.Search(term, api.SearchVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				info, err := client.RestaurantDetails(searchInfo.Restaurants[0],
					api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
					api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("RestaurantDetailsByID",
			func(ctx handler.EvalContext) (interface{}, error) {
				restID, ok := ctx.MustString("rest_id")
				if !ok {
					return nil, nil
				}
				info, err := client.RestaurantDetailsFromID(restID,
					api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
					api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("SaveRawRestaurantDetailsFromID",
			func(ctx handler.EvalContext) (interface{}, error) {
				restID, ok := ctx.MustString("rest_id")
				if !ok {
					return nil, nil
				}
				info, err := client.RawRestaurantDetailsFromID(restID,
					api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
					api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
				if err != nil {
					return nil, err
				}
				uri := fmt.Sprintf("https://www.opentable.com/r/%s", restID)
				if err := client.Cache().SaveRestaurant(ctx.Context(), uri, *info,
					api.SaveRestaurantVerbose(ctx.Bool("verbose"))); err != nil {
					return nil, err
				}
				return info, nil
			}, handler.NewHandlerCliOnly(true)),

		handler.NewHandler("SearchAndSave",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				if err := client.SearchAndSave(ctx.Context(), term, api.SearchAndSaveVerbose(ctx.Bool("verbose"))); err != nil {
					return nil, err
				}
				return nil, nil
			}, handler.NewHandlerCliOnly(true)),

		handler.NewHandler("SearchByURIAndSave",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				if err := client.SearchByURIAndSave(ctx.Context(), term, api.SearchByURIAndSaveVerbose(ctx.Bool("verbose"))); err != nil {
					return nil, err
				}
				return nil, nil
			}, handler.NewHandlerCliOnly(true)),

		handler.NewHandler("FindMatchingMenuItems",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				restID, ok := ctx.MustString("rest_id")
				if !ok {
					return nil, nil
				}
				info, err := client.RestaurantDetailsFromID(restID,
					api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
					api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
				if err != nil {
					return nil, err
				}
				items := api.FindMatchingMenuItems(info.RestaurantDetails, term)
				return items, nil
			}),

		handler.NewHandler("AllMenuItems",
			func(ctx handler.EvalContext) (interface{}, error) {
				restID, ok := ctx.MustString("rest_id")
				if !ok {
					return nil, nil
				}
				info, err := client.RestaurantDetailsFromID(restID,
					api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
					api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("FindMenuItem",
			func(ctx handler.EvalContext) (interface{}, error) {
				term, ok := ctx.MustString("term")
				if !ok {
					return nil, nil
				}
				info, err := client.FindMenuItem(term, api.FindMenuItemVerbose(ctx.Bool("verbose")))
				if err != nil {
					return nil, err
				}
				return info, nil
			}),

		handler.NewHandler("AddRestaurantsToSearchByURIs",
			func(ctx handler.EvalContext) (interface{}, error) {
				uri, ok := ctx.MustString("uri")
				if !ok {
					return nil, nil
				}
				uris := slice.NonEmptyStrings(slice.Strings(uri, ","))
				for _, uri := range uris {
					toSearch, errs := client.AddRestaurantsToSearchByURIs(ctx.Context(), uri,
						api.AddRestaurantsToSearchByURIsThreads(ctx.Int("threads")),
						api.AddRestaurantsToSearchByURIsVerbose(ctx.Bool("verbose")))
					parallel.WaitFor(
						func() {
							threads := or.Int(ctx.Int("threads"), 20)
							sleep := or.Duration(ctx.Duration("sleep"), 3*time.Second)
							var wg sync.WaitGroup
							for i := 0; i < threads; i++ {
								wg.Add(1)
								go func() {
									defer wg.Done()
									for r := range toSearch {
										if err := client.SearchRestaurantFromQueue(ctx.Context(), r, api.SearchRestaurantFromQueueVerbose(ctx.Bool("verbose"))); err != nil {
											fmt.Printf("SearchRestaurantFromQueue: %v\n", err)
										}
										if sleep > 0 {
											time.Sleep(sleep)
										}
									}
								}()
							}
							wg.Wait()
						},
						func() {
							for e := range errs {
								fmt.Printf("AddRestaurantsToSearchByURIs error: %v\n", e)
							}
						})
				}
				return nil, nil
			}, handler.NewHandlerCliOnly(true)),

		handler.NewHandler("SearchEmptyRestaurants",
			func(ctx handler.EvalContext) (interface{}, error) {
				if err := client.SearchEmptyRestaurants(ctx.Context(),
					api.SearchEmptyRestaurantsThreads(ctx.Int("threads")),
					api.SearchEmptyRestaurantsVerbose(ctx.Bool("verbose")),
					api.SearchEmptyRestaurantsSleep(ctx.Duration("sleep"))); err != nil {
					return nil, err
				}
				return nil, nil
			}, handler.NewHandlerCliOnly(true)),

		handler.NewHandler("RawListAllByURI",
			func(ctx handler.EvalContext) (interface{}, error) {
				uri, ok := ctx.MustString("uri")
				if !ok {
					return nil, nil
				}
				ress, errs := client.RawListAllByURI(uri, api.RawListAllByURIVerbose(ctx.Bool("verbose")))
				type resp struct {
					Infos  []api.RawListByURIInfo
					Errors []error
				}
				var res resp
				parallel.WaitFor(
					func() {
						for r := range ress {
							res.Infos = append(res.Infos, r)
						}
					},
					func() {
						for e := range errs {
							res.Errors = append(res.Errors, e)
						}
					})
				return ress, nil

			}),

		// space at the bottom
	}
}