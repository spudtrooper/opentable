// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RawRestaurantDetailsByURIOption struct {
	f func(*rawRestaurantDetailsByURIOptionImpl)
	s string
}

func (o RawRestaurantDetailsByURIOption) String() string { return o.s }

type RawRestaurantDetailsByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
}

func RawRestaurantDetailsByURIVerbose(verbose bool) RawRestaurantDetailsByURIOption {
	return RawRestaurantDetailsByURIOption{func(opts *rawRestaurantDetailsByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RawRestaurantDetailsByURIVerbose(bool %+v)}", verbose)}
}
func RawRestaurantDetailsByURIVerboseFlag(verbose *bool) RawRestaurantDetailsByURIOption {
	return RawRestaurantDetailsByURIOption{func(opts *rawRestaurantDetailsByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RawRestaurantDetailsByURIVerbose(bool %+v)}", verbose)}
}

func RawRestaurantDetailsByURIDebugFailures(debugFailures bool) RawRestaurantDetailsByURIOption {
	return RawRestaurantDetailsByURIOption{func(opts *rawRestaurantDetailsByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.RawRestaurantDetailsByURIDebugFailures(bool %+v)}", debugFailures)}
}
func RawRestaurantDetailsByURIDebugFailuresFlag(debugFailures *bool) RawRestaurantDetailsByURIOption {
	return RawRestaurantDetailsByURIOption{func(opts *rawRestaurantDetailsByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.RawRestaurantDetailsByURIDebugFailures(bool %+v)}", debugFailures)}
}

type rawRestaurantDetailsByURIOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *rawRestaurantDetailsByURIOptionImpl) Verbose() bool          { return r.verbose }
func (r *rawRestaurantDetailsByURIOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *rawRestaurantDetailsByURIOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *rawRestaurantDetailsByURIOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }

type RawRestaurantDetailsByURIParams struct {
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o RawRestaurantDetailsByURIParams) Options() []RawRestaurantDetailsByURIOption {
	return []RawRestaurantDetailsByURIOption{
		RawRestaurantDetailsByURIVerbose(o.Verbose),
		RawRestaurantDetailsByURIDebugFailures(o.DebugFailures),
	}
}

func makeRawRestaurantDetailsByURIOptionImpl(opts ...RawRestaurantDetailsByURIOption) *rawRestaurantDetailsByURIOptionImpl {
	res := &rawRestaurantDetailsByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRawRestaurantDetailsByURIOptions(opts ...RawRestaurantDetailsByURIOption) RawRestaurantDetailsByURIOptions {
	return makeRawRestaurantDetailsByURIOptionImpl(opts...)
}
