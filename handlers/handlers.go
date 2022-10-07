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

func CreateHandlers(client *api.Extended) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandlerFromParams("Search",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchParams)
			return client.Search(p.Term, p.Options()...)
		},
		api.SearchParams{},
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
		api.SearchAllParams{},
	)

	b.NewHandlerFromParams("LocationPicker",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.LocationPickerParams)
			if p.AuthCke != "" {
				client = client.WithAuthCke(p.AuthCke)
			}
			return client.LocationPicker(p.Options()...)
		},
		api.RestaurantsAvailabilityParams{},
	)

	b.NewHandlerFromParams("RestaurantsAvailability",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantsAvailabilityParams)
			return client.RestaurantsAvailability(p.Options()...)
		},
		api.RestaurantsAvailabilityParams{},
	)

	b.NewHandlerFromParams("SearchByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchByURIParams)
			return client.SearchByURI(p.Uri, p.Options()...)
		},
		api.SearchByURIParams{},
	)

	b.NewHandlerFromParams("RawListByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.RawListByURI(p.Uri, p.Options()...)
		},
		api.ListByURIParams{},
	)

	b.NewHandlerFromParams("ListByURI",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ListByURIParams)
			return client.ListByURI(p.Uri, p.Options()...)
		},
		api.ListByURIParams{},
	)

	b.NewHandlerFromParams("RestaurantDetailsFromSearch",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromSearchParams)
			return client.RestaurantDetailsFromSearch(p.Term, p.Options()...)
		},
		api.RestaurantDetailsFromSearchParams{},
	)

	b.NewHandlerFromParams("RestaurantDetailsByID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromIDParams)
			return client.RestaurantDetailsFromID(p.Id, p.Options()...)
		},
		api.RestaurantDetailsFromIDParams{},
	)

	b.NewHandlerFromParams("SaveRawRestaurantDetailsFromID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SaveRawRestaurantDetailsFromIDParams)
			return client.SaveRawRestaurantDetailsFromID(ctx, p.RestID, p.Options()...)
		},
		api.SaveRawRestaurantDetailsFromIDParams{},
		handler.NewHandlerMethod("POST"),
	)

	b.NewHandlerFromParams("SearchAndSave",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchAndSaveParams)
			return nil, client.SearchAndSave(ctx, p.Term, p.Options()...)
		},
		api.SearchAndSaveParams{},
		handler.NewHandlerMethod("POST"),
	)

	b.NewHandlerFromParams("SearchByURIAndSave",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchByURIAndSaveParams)
			return nil, client.SearchByURIAndSave(ctx, p.Uri, p.Options()...)
		},
		api.SearchByURIAndSaveParams{},
	)

	b.NewHandlerFromParams("FindMatchingMenuItemsFromRestaurantID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.FindMatchingMenuItemsFromRestaurantIDParams)
			return client.FindMatchingMenuItemsFromRestaurantID(p.RestID, p.Term, p.Options()...)
		},
		api.FindMatchingMenuItemsFromRestaurantIDParams{},
	)

	b.NewHandlerFromParams("RestaurantDetailsFromID",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RestaurantDetailsFromIDParams)
			return client.RestaurantDetailsFromID(p.Id, p.Options()...)
		},
		api.RestaurantDetailsFromIDParams{},
	)

	b.NewHandlerFromParams("FindMenuItem",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.FindMenuItemParams)
			return client.FindMenuItem(p.Term, p.Options()...)
		},
		api.FindMenuItemParams{},
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
		api.SearchEmptyRestaurantsParams{},
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
