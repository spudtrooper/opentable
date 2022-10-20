// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RestaurantDetailsByURIOption struct {
	f func(*restaurantDetailsByURIOptionImpl)
	s string
}

func (o RestaurantDetailsByURIOption) String() string { return o.s }

type RestaurantDetailsByURIOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func RestaurantDetailsByURIDebugFailures(debugFailures bool) RestaurantDetailsByURIOption {
	return RestaurantDetailsByURIOption{func(opts *restaurantDetailsByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsByURIDebugFailures(bool %+v)}", debugFailures)}
}
func RestaurantDetailsByURIDebugFailuresFlag(debugFailures *bool) RestaurantDetailsByURIOption {
	return RestaurantDetailsByURIOption{func(opts *restaurantDetailsByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsByURIDebugFailures(bool %+v)}", debugFailures)}
}

func RestaurantDetailsByURIVerbose(verbose bool) RestaurantDetailsByURIOption {
	return RestaurantDetailsByURIOption{func(opts *restaurantDetailsByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RestaurantDetailsByURIVerbose(bool %+v)}", verbose)}
}
func RestaurantDetailsByURIVerboseFlag(verbose *bool) RestaurantDetailsByURIOption {
	return RestaurantDetailsByURIOption{func(opts *restaurantDetailsByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RestaurantDetailsByURIVerbose(bool %+v)}", verbose)}
}

type restaurantDetailsByURIOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *restaurantDetailsByURIOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsByURIOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }
func (r *restaurantDetailsByURIOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsByURIOptionImpl) HasVerbose() bool       { return r.has_verbose }

type RestaurantDetailsByURIParams struct {
	DebugFailures bool   `json:"debug_failures"`
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o RestaurantDetailsByURIParams) Options() []RestaurantDetailsByURIOption {
	return []RestaurantDetailsByURIOption{
		RestaurantDetailsByURIDebugFailures(o.DebugFailures),
		RestaurantDetailsByURIVerbose(o.Verbose),
	}
}

func makeRestaurantDetailsByURIOptionImpl(opts ...RestaurantDetailsByURIOption) *restaurantDetailsByURIOptionImpl {
	res := &restaurantDetailsByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRestaurantDetailsByURIOptions(opts ...RestaurantDetailsByURIOption) RestaurantDetailsByURIOptions {
	return makeRestaurantDetailsByURIOptionImpl(opts...)
}
