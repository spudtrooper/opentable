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
		func(ip any) (any, error) {
			p := ip.(api.SearchParams)
			return client.Search(p.Term, p.Options()...)
		},
		func() any { return api.SearchParams{} },
	)

	b.NewHandlerFromParams("SearchAll",
		func(ip any) (any, error) {
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
		func(ip any) (any, error) {
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
		func(ip any) (any, error) {
			p := ip.(api.RestaurantsAvailabilityParams)
			return client.RestaurantsAvailability(p.Options()...)
		},
		func() any { return api.RestaurantsAvailabilityParams{} },
	)

	b.NewHandlerFromParams("SearchByURI",
		func(ip any) (any, error) {
			p := ip.(api.SearchByURIParams)
			return client.SearchByURI(p.Uri, p.Options()...)
		},
		func() any { return api.SearchByURIParams{} },
	)

	b.NewHandlerFromParams("RawListByURI",
		func(ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.RawListByURI(p.Uri, p.Options()...)
		},
		func() any { return api.ListByURIParams{} },
	)

	b.NewHandlerFromParams("ListByURI",
		func(ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.ListByURI(p.Uri, p.Options()...)
		},
		func() any { return api.ListByURIParams{} },
	)

	b.NewHandler("RestaurantDetails",
		func(ctx handler.EvalContext) (interface{}, error) {
			term, ok := ctx.MustString("term")
			if !ok {
				return nil, nil
			}
			searchInfo, err := client.Search(term, api.SearchVerbose(ctx.Bool("verbose")))
			if err != nil {
				return nil, err
			}
			return client.RestaurantDetails(searchInfo.Restaurants[0],
				api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
				api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
		},
		handler.Params().
			RequiredString("term").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandler("RestaurantDetailsByID",
		func(ctx handler.EvalContext) (interface{}, error) {
			restID, ok := ctx.MustString("rest_id")
			if !ok {
				return nil, nil
			}
			return client.RestaurantDetailsFromID(restID,
				api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
				api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
		},
		handler.Params().
			RequiredString("rest_id").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandler("SaveRawRestaurantDetailsFromID",
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
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			RequiredString("rest_id").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandler("SearchAndSave",
		func(ctx handler.EvalContext) (interface{}, error) {
			term, ok := ctx.MustString("term")
			if !ok {
				return nil, nil
			}
			if err := client.SearchAndSave(ctx.Context(), term, api.SearchAndSaveVerbose(ctx.Bool("verbose"))); err != nil {
				return nil, err
			}
			return nil, nil
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			RequiredString("term").
			Bool("verbose").
			BuildOption())

	b.NewHandler("SearchByURIAndSave",
		func(ctx handler.EvalContext) (interface{}, error) {
			term, ok := ctx.MustString("term")
			if !ok {
				return nil, nil
			}
			if err := client.SearchByURIAndSave(ctx.Context(), term, api.SearchByURIAndSaveVerbose(ctx.Bool("verbose"))); err != nil {
				return nil, err
			}
			return nil, nil
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			RequiredString("term").
			Bool("verbose").
			BuildOption())

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
				api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
				api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
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

	b.NewHandler("AllMenuItems",
		func(ctx handler.EvalContext) (interface{}, error) {
			restID, ok := ctx.MustString("rest_id")
			if !ok {
				return nil, nil
			}
			return client.RestaurantDetailsFromID(restID,
				api.RestaurantDetailsVerbose(ctx.Bool("verbose")),
				api.RestaurantDetailsDebugFailures(ctx.Bool("debug_failures")))
		},
		handler.Params().
			RequiredString("rest_id").
			Bool("verbose").
			Bool("debug_failures").
			BuildOption())

	b.NewHandler("FindMenuItem",
		func(ctx handler.EvalContext) (interface{}, error) {
			term, ok := ctx.MustString("term")
			if !ok {
				return nil, nil
			}
			return client.FindMenuItem(term,
				api.FindMenuItemVerbose(ctx.Bool("verbose")),
				api.FindMenuItemLongitude(ctx.Float32("longitude")),
				api.FindMenuItemLatitude(ctx.Float32("latitude")),
				api.FindMenuItemMetroID(ctx.Int("metro_id")),
			)
		},
		handler.Params().
			RequiredString("term").
			Bool("verbose").
			Float32("latitude").
			Float32("longitude").
			Int("metro_id").
			BuildOption())

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

	b.NewHandler("SearchEmptyRestaurants",
		func(ctx handler.EvalContext) (interface{}, error) {
			if err := client.SearchEmptyRestaurants(ctx.Context(),
				api.SearchEmptyRestaurantsThreads(ctx.Int("threads")),
				api.SearchEmptyRestaurantsVerbose(ctx.Bool("verbose")),
				api.SearchEmptyRestaurantsSleep(ctx.Duration("sleep"))); err != nil {
				return nil, err
			}
			return nil, nil
		},
		handler.NewHandlerMethod("POST"),
		handler.Params().
			Bool("verbose").
			Int("threads").
			Duration("sleep").
			BuildOption())

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
