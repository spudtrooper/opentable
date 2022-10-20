// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RestaurantDetailsOption struct {
	f func(*restaurantDetailsOptionImpl)
	s string
}

func (o RestaurantDetailsOption) String() string { return o.s }

type RestaurantDetailsOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func RestaurantDetailsDebugFailures(debugFailures bool) RestaurantDetailsOption {
	return RestaurantDetailsOption{func(opts *restaurantDetailsOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsDebugFailures(bool %+v)}", debugFailures)}
}
func RestaurantDetailsDebugFailuresFlag(debugFailures *bool) RestaurantDetailsOption {
	return RestaurantDetailsOption{func(opts *restaurantDetailsOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsDebugFailures(bool %+v)}", debugFailures)}
}

func RestaurantDetailsVerbose(verbose bool) RestaurantDetailsOption {
	return RestaurantDetailsOption{func(opts *restaurantDetailsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RestaurantDetailsVerbose(bool %+v)}", verbose)}
}
func RestaurantDetailsVerboseFlag(verbose *bool) RestaurantDetailsOption {
	return RestaurantDetailsOption{func(opts *restaurantDetailsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RestaurantDetailsVerbose(bool %+v)}", verbose)}
}

type restaurantDetailsOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *restaurantDetailsOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }
func (r *restaurantDetailsOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsOptionImpl) HasVerbose() bool       { return r.has_verbose }

type RestaurantDetailsParams struct {
	DebugFailures bool `json:"debug_failures"`
	Verbose       bool `json:"verbose"`
}

func (o RestaurantDetailsParams) Options() []RestaurantDetailsOption {
	return []RestaurantDetailsOption{
		RestaurantDetailsDebugFailures(o.DebugFailures),
		RestaurantDetailsVerbose(o.Verbose),
	}
}

func makeRestaurantDetailsOptionImpl(opts ...RestaurantDetailsOption) *restaurantDetailsOptionImpl {
	res := &restaurantDetailsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRestaurantDetailsOptions(opts ...RestaurantDetailsOption) RestaurantDetailsOptions {
	return makeRestaurantDetailsOptionImpl(opts...)
}
