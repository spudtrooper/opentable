// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/parallel"
	"github.com/spudtrooper/goutil/slice"
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opentable/api"
)

type LocationPickerParams struct {
	AuthCke  string `json:"auth_cke"`
	Tld      string `json:"tld"`
	MetroID  int    `json:"metro_id"`
	DomainID int    `json:"domain_id"`
	Verbose  bool   `json:"verbose"`
}

func CreateHandlers(client *api.Extended) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandlerFromParams("Search",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchParams)
			return client.Search(p.Term, p.Options()...)
		},
		func() any { return api.SearchParams{} },
	)

	b.NewHandlerFromParams("SearchAll",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchAllParams)
			infos, errs := client.SearchAll(p.Term, p.Options()...)
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
		},
		func() any { return api.SearchAllParams{} },
	)

	b.NewHandlerFromParams("LocationPicker",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(LocationPickerParams)
			if p.AuthCke != "" {
				client = client.WithAuthCke(p.AuthCke)
			}
			return client.LocationPicker(
				api.LocationPickerVerbose(p.Verbose),
				api.LocationPickerTld(p.Tld),
				api.LocationPickerMetroID(p.MetroID),
				api.LocationPickerDomainID(p.DomainID),
			)
		},
		func() any { return LocationPickerParams{} },
	)

	b.NewHandlerFromParams("RestaurantsAvailability",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantsAvailabilityParams)
			return client.RestaurantsAvailability(p.Options()...)
		},
		func() any { return api.RestaurantsAvailabilityParams{} },
	)

	b.NewHandlerFromParams("SearchByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchByURIParams)
			return client.SearchByURI(p.Uri, p.Options()...)
		},
		func() any { return api.SearchByURIParams{} },
	)

	b.NewHandlerFromParams("RawListByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.RawListByURI(p.Uri, p.Options()...)
		},
		func() any { return api.ListByURIParams{} },
	)

	b.NewHandlerFromParams("ListByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.ListByURI(p.Uri, p.Options()...)
		},
		func() any { return api.ListByURIParams{} },
	)

	b.NewHandlerFromParams("RestaurantDetailsFromSearch",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromSearchParams)
			return client.RestaurantDetailsFromSearch(p.Term, p.Options()...)
		},
		func() any { return api.RestaurantDetailsFromSearchParams{} },
	)

	b.NewHandlerFromParams("RestaurantDetailsByID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromIDParams)
			return client.RestaurantDetailsFromID(p.Id, p.Options()...)
		},
		func() any { return api.RestaurantDetailsFromIDParams{} },
	)

	b.NewHandler("SaveRawRestaurantDetailsFromID",
		func(ctx handler.EvalContext) (interface{}, error) {
			restID, ok := ctx.MustString("rest_id")
			if !ok {
				return nil, nil
			}
			info, err := client.RawRestaurantDetailsFromID(restID,
				api.RawRestaurantDetailsFromIDVerbose(ctx.Bool("verbose")),
				api.RawRestaurantDetailsFromIDDebugFailures(ctx.Bool("debug_failures")))
			if err != nil {
				return nil, err
			}
			uri := fmt.Sprintf("https://www.opentable.com/r/%s", restID)
			if err := client.Cache().SaveRestaurant(ctx.Context(), uri, *info,
				api.SaveRestaurantVerbose(ctx.Bool("verbose"))); err != nil {
				return nil, err
			}
			return info, nil
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			RequiredString("rest_id").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandlerFromParams("SearchAndSave",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchAndSaveParams)
			return nil, client.SearchAndSave(ctx, p.Term, p.Options()...)
		},
		func() any { return api.SearchAndSaveParams{} },
	)

	b.NewHandlerFromParams("SearchByURIAndSave",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchByURIAndSaveParams)
			return nil, client.SearchByURIAndSave(ctx, p.Uri, p.Options()...)
		},
		func() any { return api.SearchByURIAndSaveParams{} },
	)

	b.NewHandler("FindMatchingMenuItems",
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
				api.RestaurantDetailsFromIDVerbose(ctx.Bool("verbose")),
				api.RestaurantDetailsFromIDDebugFailures(ctx.Bool("debug_failures")))
			if err != nil {
				return nil, err
			}
			items := api.FindMatchingMenuItems(info.RestaurantDetails, term)
			return items, nil
		},
		handler.Params().
			RequiredString("term").
			RequiredString("rest_id").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandlerFromParams("RestaurantDetailsFromID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromIDParams)
			return client.RestaurantDetailsFromID(p.Id, p.Options()...)
		},
		func() any { return api.RestaurantDetailsFromIDParams{} },
	)

	b.NewHandlerFromParams("FindMenuItem",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.FindMenuItemParams)
			return client.FindMenuItem(p.Term, p.Options()...)
		},
		func() any { return api.FindMenuItemParams{} },
	)

	b.NewHandler("AddRestaurantsToSearchByURIs",
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
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			RequiredString("uri").
			Bool("verbose").
			Int("threads").
			Duration("sleep").
			BuildOption())

	b.NewHandlerFromParams("SearchEmptyRestaurants",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchEmptyRestaurantsParams)
			return nil, client.SearchEmptyRestaurants(ctx, p.Options()...)
		},
		func() any { return api.SearchEmptyRestaurantsParams{} },
	)

	b.NewHandler("RawListAllByURI",
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

		},
		handler.Params().
			RequiredString("uri").
			Bool("verbose").
			BuildOption())

	return b.Build()
}
