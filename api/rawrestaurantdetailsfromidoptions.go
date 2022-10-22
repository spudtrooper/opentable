// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RawRestaurantDetailsFromIDOption struct {
	f func(*rawRestaurantDetailsFromIDOptionImpl)
	s string
}

func (o RawRestaurantDetailsFromIDOption) String() string { return o.s }

type RawRestaurantDetailsFromIDOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func RawRestaurantDetailsFromIDDebugFailures(debugFailures bool) RawRestaurantDetailsFromIDOption {
	return RawRestaurantDetailsFromIDOption{func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.RawRestaurantDetailsFromIDDebugFailures(bool %+v)", debugFailures)}
}
func RawRestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) RawRestaurantDetailsFromIDOption {
	return RawRestaurantDetailsFromIDOption{func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.RawRestaurantDetailsFromIDDebugFailures(bool %+v)", debugFailures)}
}

func RawRestaurantDetailsFromIDVerbose(verbose bool) RawRestaurantDetailsFromIDOption {
	return RawRestaurantDetailsFromIDOption{func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RawRestaurantDetailsFromIDVerbose(bool %+v)", verbose)}
}
func RawRestaurantDetailsFromIDVerboseFlag(verbose *bool) RawRestaurantDetailsFromIDOption {
	return RawRestaurantDetailsFromIDOption{func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RawRestaurantDetailsFromIDVerbose(bool %+v)", verbose)}
}

type rawRestaurantDetailsFromIDOptionImpl struct {
	debugFailures     bool
	has_debugFailures bool
	verbose           bool
	has_verbose       bool
}

func (r *rawRestaurantDetailsFromIDOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *rawRestaurantDetailsFromIDOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }
func (r *rawRestaurantDetailsFromIDOptionImpl) Verbose() bool          { return r.verbose }
func (r *rawRestaurantDetailsFromIDOptionImpl) HasVerbose() bool       { return r.has_verbose }

type RawRestaurantDetailsFromIDParams struct {
	DebugFailures bool   `json:"debug_failures"`
	Id            string `json:"id" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o RawRestaurantDetailsFromIDParams) Options() []RawRestaurantDetailsFromIDOption {
	return []RawRestaurantDetailsFromIDOption{
		RawRestaurantDetailsFromIDDebugFailures(o.DebugFailures),
		RawRestaurantDetailsFromIDVerbose(o.Verbose),
	}
}

func makeRawRestaurantDetailsFromIDOptionImpl(opts ...RawRestaurantDetailsFromIDOption) *rawRestaurantDetailsFromIDOptionImpl {
	res := &rawRestaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRawRestaurantDetailsFromIDOptions(opts ...RawRestaurantDetailsFromIDOption) RawRestaurantDetailsFromIDOptions {
	return makeRawRestaurantDetailsFromIDOptionImpl(opts...)
}
