// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RawSearchAllByURIOption func(*rawSearchAllByURIOptionImpl)

type RawSearchAllByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	StartPage() int
	HasStartPage() bool
}

func RawSearchAllByURIVerbose(verbose bool) RawSearchAllByURIOption {
	return func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RawSearchAllByURIVerboseFlag(verbose *bool) RawSearchAllByURIOption {
	return func(opts *rawSearchAllByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RawSearchAllByURIStartPage(startPage int) RawSearchAllByURIOption {
	return func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}
}
func RawSearchAllByURIStartPageFlag(startPage *int) RawSearchAllByURIOption {
	return func(opts *rawSearchAllByURIOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}
}

type rawSearchAllByURIOptionImpl struct {
	verbose       bool
	has_verbose   bool
	startPage     int
	has_startPage bool
}

func (r *rawSearchAllByURIOptionImpl) Verbose() bool      { return r.verbose }
func (r *rawSearchAllByURIOptionImpl) HasVerbose() bool   { return r.has_verbose }
func (r *rawSearchAllByURIOptionImpl) StartPage() int     { return r.startPage }
func (r *rawSearchAllByURIOptionImpl) HasStartPage() bool { return r.has_startPage }

func makeRawSearchAllByURIOptionImpl(opts ...RawSearchAllByURIOption) *rawSearchAllByURIOptionImpl {
	res := &rawSearchAllByURIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRawSearchAllByURIOptions(opts ...RawSearchAllByURIOption) RawSearchAllByURIOptions {
	return makeRawSearchAllByURIOptionImpl(opts...)
}
