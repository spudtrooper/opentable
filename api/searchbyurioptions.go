// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SearchByURIOption func(*searchByURIOptionImpl)

type SearchByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
}

func SearchByURIVerbose(verbose bool) SearchByURIOption {
	return func(opts *searchByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchByURIVerboseFlag(verbose *bool) SearchByURIOption {
	return func(opts *searchByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func SearchByURIDebugFailures(debugFailures bool) SearchByURIOption {
	return func(opts *searchByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func SearchByURIDebugFailuresFlag(debugFailures *bool) SearchByURIOption {
	return func(opts *searchByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

type searchByURIOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (s *searchByURIOptionImpl) Verbose() bool          { return s.verbose }
func (s *searchByURIOptionImpl) HasVerbose() bool       { return s.has_verbose }
func (s *searchByURIOptionImpl) DebugFailures() bool    { return s.debugFailures }
func (s *searchByURIOptionImpl) HasDebugFailures() bool { return s.has_debugFailures }

func makeSearchByURIOptionImpl(opts ...SearchByURIOption) *searchByURIOptionImpl {
	res := &searchByURIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchByURIOptions(opts ...SearchByURIOption) SearchByURIOptions {
	return makeSearchByURIOptionImpl(opts...)
}