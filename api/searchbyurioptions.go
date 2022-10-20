// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SearchByURIOption struct {
	f func(*searchByURIOptionImpl)
	s string
}

func (o SearchByURIOption) String() string { return o.s }

type SearchByURIOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func SearchByURIDebugFailures(debugFailures bool) SearchByURIOption {
	return SearchByURIOption{func(opts *searchByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.SearchByURIDebugFailures(bool %+v)}", debugFailures)}
}
func SearchByURIDebugFailuresFlag(debugFailures *bool) SearchByURIOption {
	return SearchByURIOption{func(opts *searchByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.SearchByURIDebugFailures(bool %+v)}", debugFailures)}
}

func SearchByURIVerbose(verbose bool) SearchByURIOption {
	return SearchByURIOption{func(opts *searchByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchByURIVerbose(bool %+v)}", verbose)}
}
func SearchByURIVerboseFlag(verbose *bool) SearchByURIOption {
	return SearchByURIOption{func(opts *searchByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchByURIVerbose(bool %+v)}", verbose)}
}

type searchByURIOptionImpl struct {
	debugFailures     bool
	has_debugFailures bool
	verbose           bool
	has_verbose       bool
}

func (s *searchByURIOptionImpl) DebugFailures() bool    { return s.debugFailures }
func (s *searchByURIOptionImpl) HasDebugFailures() bool { return s.has_debugFailures }
func (s *searchByURIOptionImpl) Verbose() bool          { return s.verbose }
func (s *searchByURIOptionImpl) HasVerbose() bool       { return s.has_verbose }

type SearchByURIParams struct {
	DebugFailures bool   `json:"debug_failures"`
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o SearchByURIParams) Options() []SearchByURIOption {
	return []SearchByURIOption{
		SearchByURIDebugFailures(o.DebugFailures),
		SearchByURIVerbose(o.Verbose),
	}
}

func makeSearchByURIOptionImpl(opts ...SearchByURIOption) *searchByURIOptionImpl {
	res := &searchByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchByURIOptions(opts ...SearchByURIOption) SearchByURIOptions {
	return makeSearchByURIOptionImpl(opts...)
}
