// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type ListByURIOption struct {
	f func(*listByURIOptionImpl)
	s string
}

func (o ListByURIOption) String() string { return o.s }

type ListByURIOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	Verbose() bool
	HasVerbose() bool
}

func ListByURIDebugFailures(debugFailures bool) ListByURIOption {
	return ListByURIOption{func(opts *listByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.ListByURIDebugFailures(bool %+v)}", debugFailures)}
}
func ListByURIDebugFailuresFlag(debugFailures *bool) ListByURIOption {
	return ListByURIOption{func(opts *listByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.ListByURIDebugFailures(bool %+v)}", debugFailures)}
}

func ListByURIVerbose(verbose bool) ListByURIOption {
	return ListByURIOption{func(opts *listByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.ListByURIVerbose(bool %+v)}", verbose)}
}
func ListByURIVerboseFlag(verbose *bool) ListByURIOption {
	return ListByURIOption{func(opts *listByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.ListByURIVerbose(bool %+v)}", verbose)}
}

type listByURIOptionImpl struct {
	debugFailures     bool
	has_debugFailures bool
	verbose           bool
	has_verbose       bool
}

func (l *listByURIOptionImpl) DebugFailures() bool    { return l.debugFailures }
func (l *listByURIOptionImpl) HasDebugFailures() bool { return l.has_debugFailures }
func (l *listByURIOptionImpl) Verbose() bool          { return l.verbose }
func (l *listByURIOptionImpl) HasVerbose() bool       { return l.has_verbose }

type ListByURIParams struct {
	DebugFailures bool   `json:"debug_failures"`
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o ListByURIParams) Options() []ListByURIOption {
	return []ListByURIOption{
		ListByURIDebugFailures(o.DebugFailures),
		ListByURIVerbose(o.Verbose),
	}
}

func makeListByURIOptionImpl(opts ...ListByURIOption) *listByURIOptionImpl {
	res := &listByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeListByURIOptions(opts ...ListByURIOption) ListByURIOptions {
	return makeListByURIOptionImpl(opts...)
}
