// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type ListByURIOption struct {
	f func(*listByURIOptionImpl)
	s string
}

func (o ListByURIOption) String() string { return o.s }

type ListByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
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

type listByURIOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (l *listByURIOptionImpl) Verbose() bool          { return l.verbose }
func (l *listByURIOptionImpl) HasVerbose() bool       { return l.has_verbose }
func (l *listByURIOptionImpl) DebugFailures() bool    { return l.debugFailures }
func (l *listByURIOptionImpl) HasDebugFailures() bool { return l.has_debugFailures }

type ListByURIParams struct {
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o ListByURIParams) Options() []ListByURIOption {
	return []ListByURIOption{
		ListByURIVerbose(o.Verbose),
		ListByURIDebugFailures(o.DebugFailures),
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
