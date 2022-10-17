// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type FindMatchingMenuItemsFromRestaurantIDOption struct {
	f func(*findMatchingMenuItemsFromRestaurantIDOptionImpl)
	s string
}

func (o FindMatchingMenuItemsFromRestaurantIDOption) String() string { return o.s }

type FindMatchingMenuItemsFromRestaurantIDOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
	ToRestaurantDetailsFromIDOptions() []RestaurantDetailsFromIDOption
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

type findMatchingMenuItemsFromRestaurantIDOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) Verbose() bool    { return f.verbose }
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) HasVerbose() bool { return f.has_verbose }
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) DebugFailures() bool {
	return f.debugFailures
}
func (f *findMatchingMenuItemsFromRestaurantIDOptionImpl) HasDebugFailures() bool {
	return f.has_debugFailures
}

type FindMatchingMenuItemsFromRestaurantIDParams struct {
	RestID        string `json:"rest_id" required:"true"`
	Term          string `json:"term" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o FindMatchingMenuItemsFromRestaurantIDParams) Options() []FindMatchingMenuItemsFromRestaurantIDOption {
	return []FindMatchingMenuItemsFromRestaurantIDOption{
		FindMatchingMenuItemsFromRestaurantIDVerbose(o.Verbose),
		FindMatchingMenuItemsFromRestaurantIDDebugFailures(o.DebugFailures),
	}
}

// ToRestaurantDetailsFromIDOptions converts FindMatchingMenuItemsFromRestaurantIDOption to an array of RestaurantDetailsFromIDOption
func (o *findMatchingMenuItemsFromRestaurantIDOptionImpl) ToRestaurantDetailsFromIDOptions() []RestaurantDetailsFromIDOption {
	return []RestaurantDetailsFromIDOption{
		RestaurantDetailsFromIDVerbose(o.Verbose()),
		RestaurantDetailsFromIDDebugFailures(o.DebugFailures()),
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
