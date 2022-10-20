// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RestaurantDetailsFromIDOption struct {
	f func(*restaurantDetailsFromIDOptionImpl)
	s string
}

func (o RestaurantDetailsFromIDOption) String() string { return o.s }

type RestaurantDetailsFromIDOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func RestaurantDetailsFromIDDebugFailures(debugFailures bool) RestaurantDetailsFromIDOption {
	return RestaurantDetailsFromIDOption{func(opts *restaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsFromIDDebugFailures(bool %+v)}", debugFailures)}
}
func RestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) RestaurantDetailsFromIDOption {
	return RestaurantDetailsFromIDOption{func(opts *restaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.RestaurantDetailsFromIDDebugFailures(bool %+v)}", debugFailures)}
}

func RestaurantDetailsFromIDVerbose(verbose bool) RestaurantDetailsFromIDOption {
	return RestaurantDetailsFromIDOption{func(opts *restaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RestaurantDetailsFromIDVerbose(bool %+v)}", verbose)}
}
func RestaurantDetailsFromIDVerboseFlag(verbose *bool) RestaurantDetailsFromIDOption {
	return RestaurantDetailsFromIDOption{func(opts *restaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RestaurantDetailsFromIDVerbose(bool %+v)}", verbose)}
}

type restaurantDetailsFromIDOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *restaurantDetailsFromIDOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsFromIDOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }
func (r *restaurantDetailsFromIDOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsFromIDOptionImpl) HasVerbose() bool       { return r.has_verbose }

type RestaurantDetailsFromIDParams struct {
	DebugFailures bool   `json:"debug_failures"`
	Id            string `json:"id" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o RestaurantDetailsFromIDParams) Options() []RestaurantDetailsFromIDOption {
	return []RestaurantDetailsFromIDOption{
		RestaurantDetailsFromIDDebugFailures(o.DebugFailures),
		RestaurantDetailsFromIDVerbose(o.Verbose),
	}
}

func makeRestaurantDetailsFromIDOptionImpl(opts ...RestaurantDetailsFromIDOption) *restaurantDetailsFromIDOptionImpl {
	res := &restaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRestaurantDetailsFromIDOptions(opts ...RestaurantDetailsFromIDOption) RestaurantDetailsFromIDOptions {
	return makeRestaurantDetailsFromIDOptionImpl(opts...)
}
