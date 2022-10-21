// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type FindMatchingMenuItemsFromRestaurantIDOption struct {
	f func(*findMatchingMenuItemsFromRestaurantIDOptionImpl)
	s string
}

func (o FindMatchingMenuItemsFromRestaurantIDOption) String() string { return o.s }

type FindMatchingMenuItemsFromRestaurantIDOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
	ToRestaurantDetailsFromIDOptions() []RestaurantDetailsFromIDOption
}

func FindMatchingMenuItemsFromRestaurantIDDebugFailures(debugFailures bool) FindMatchingMenuItemsFromRestaurantIDOption {
	return FindMatchingMenuItemsFromRestaurantIDOption{func(opts *findMatchingMenuItemsFromRestaurantIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.FindMatchingMenuItemsFromRestaurantIDDebugFailures(bool %+v)}", debugFailures)}
}
func FindMatchingMenuItemsFromRestaurantIDDebugFailuresFlag(debugFailures *bool) FindMatchingMenuItemsFromRestaurantIDOption {
	return FindMatchingMenuItemsFromRestaurantIDOption{func(opts *findMatchingMenuItemsFromRestaurantIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.FindMatchingMenuItemsFromRestaurantIDDebugFailures(bool %+v)}", debugFailures)}
}

func FindMatchingMenuItemsFromRestaurantIDVerbose(verbose bool) FindMatchingMenuItemsFromRestaurantIDOption {
	return FindMatchingMenuItemsFromRestaurantIDOption{func(opts *findMatchingMenuItemsFromRestaurantIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.FindMatchingMenuItemsFromRestaurantIDVerbose(bool %+v)}", verbose)}
}
func FindMatchingMenuItemsFromRestaurantIDVerboseFlag(verbose *bool) FindMatchingMenuItemsFromRestaurantIDOption {
	return FindMatchingMenuItemsFromRestaurantIDOption{func(opts *findMatchingMenuItemsFromRestaurantIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.FindMatchingMenuItemsFromRestaurantIDVerbose(bool %+v)}", verbose)}
}

type findMatchingMenuItemsFromRestaurantIDOptionImpl struct {
	debugFailures     bool
	has_debugFailures bool
	verbose           bool
	has_verbose       bool
}

func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) DebugFailures() bool {
	return f.debugFailures
}
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) HasDebugFailures() bool {
	return f.has_debugFailures
}
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) Verbose() bool    { return f.verbose }
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) HasVerbose() bool { return f.has_verbose }

type FindMatchingMenuItemsFromRestaurantIDParams struct {
	DebugFailures bool   `json:"debug_failures"`
	RestID        string `json:"rest_id" required:"true"`
	Term          string `json:"term" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o FindMatchingMenuItemsFromRestaurantIDParams) Options() []FindMatchingMenuItemsFromRestaurantIDOption {
	return []FindMatchingMenuItemsFromRestaurantIDOption{
		FindMatchingMenuItemsFromRestaurantIDDebugFailures(o.DebugFailures),
		FindMatchingMenuItemsFromRestaurantIDVerbose(o.Verbose),
	}
}

// ToRestaurantDetailsFromIDOptions converts FindMatchingMenuItemsFromRestaurantIDOption to an array of RestaurantDetailsFromIDOption
func (o *findMatchingMenuItemsFromRestaurantIDOptionImpl) ToRestaurantDetailsFromIDOptions() []RestaurantDetailsFromIDOption {
	return []RestaurantDetailsFromIDOption{
		RestaurantDetailsFromIDDebugFailures(o.DebugFailures()),
		RestaurantDetailsFromIDVerbose(o.Verbose()),
	}
}

func makeFindMatchingMenuItemsFromRestaurantIDOptionImpl(opts ...FindMatchingMenuItemsFromRestaurantIDOption) *findMatchingMenuItemsFromRestaurantIDOptionImpl {
	res := &findMatchingMenuItemsFromRestaurantIDOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFindMatchingMenuItemsFromRestaurantIDOptions(opts ...FindMatchingMenuItemsFromRestaurantIDOption) FindMatchingMenuItemsFromRestaurantIDOptions {
	return makeFindMatchingMenuItemsFromRestaurantIDOptionImpl(opts...)
}
