// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RawSearchAllByURIOption struct {
	f func(*rawSearchAllByURIOptionImpl)
	s string
}

func (o RawSearchAllByURIOption) String() string { return o.s }

type RawSearchAllByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	StartPage() int
	HasStartPage() bool
	Threads() int
	HasThreads() bool
	Sync() bool
	HasSync() bool
}

func RawSearchAllByURIVerbose(verbose bool) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RawSearchAllByURIVerbose(bool %+v)}", verbose)}
}
func RawSearchAllByURIVerboseFlag(verbose *bool) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RawSearchAllByURIVerbose(bool %+v)}", verbose)}
}

func RawSearchAllByURIStartPage(startPage int) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}, fmt.Sprintf("api.RawSearchAllByURIStartPage(int %+v)}", startPage)}
}
func RawSearchAllByURIStartPageFlag(startPage *int) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}, fmt.Sprintf("api.RawSearchAllByURIStartPage(int %+v)}", startPage)}
}

func RawSearchAllByURIThreads(threads int) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.RawSearchAllByURIThreads(int %+v)}", threads)}
}
func RawSearchAllByURIThreadsFlag(threads *int) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.RawSearchAllByURIThreads(int %+v)}", threads)}
}

func RawSearchAllByURISync(sync bool) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		opts.has_sync = true
		opts.sync = sync
	}, fmt.Sprintf("api.RawSearchAllByURISync(bool %+v)}", sync)}
}
func RawSearchAllByURISyncFlag(sync *bool) RawSearchAllByURIOption {
	return RawSearchAllByURIOption{func(opts *rawSearchAllByURIOptionImpl) {
		if sync == nil {
			return
		}
		opts.has_sync = true
		opts.sync = *sync
	}, fmt.Sprintf("api.RawSearchAllByURISync(bool %+v)}", sync)}
}

type rawSearchAllByURIOptionImpl struct {
	verbose       bool
	has_verbose   bool
	startPage     int
	has_startPage bool
	threads       int
	has_threads   bool
	sync          bool
	has_sync      bool
}

func (r *rawSearchAllByURIOptionImpl) Verbose() bool      { return r.verbose }
func (r *rawSearchAllByURIOptionImpl) HasVerbose() bool   { return r.has_verbose }
func (r *rawSearchAllByURIOptionImpl) StartPage() int     { return r.startPage }
func (r *rawSearchAllByURIOptionImpl) HasStartPage() bool { return r.has_startPage }
func (r *rawSearchAllByURIOptionImpl) Threads() int       { return r.threads }
func (r *rawSearchAllByURIOptionImpl) HasThreads() bool   { return r.has_threads }
func (r *rawSearchAllByURIOptionImpl) Sync() bool         { return r.sync }
func (r *rawSearchAllByURIOptionImpl) HasSync() bool      { return r.has_sync }

type RawSearchAllByURIParams struct {
	Uri       string `json:"uri" required:"true"`
	Verbose   bool   `json:"verbose"`
	StartPage int    `json:"start_page"`
	Threads   int    `json:"threads"`
	Sync      bool   `json:"sync"`
}

func (o RawSearchAllByURIParams) Options() []RawSearchAllByURIOption {
	return []RawSearchAllByURIOption{
		RawSearchAllByURIVerbose(o.Verbose),
		RawSearchAllByURIStartPage(o.StartPage),
		RawSearchAllByURIThreads(o.Threads),
		RawSearchAllByURISync(o.Sync),
	}
}

func makeRawSearchAllByURIOptionImpl(opts ...RawSearchAllByURIOption) *rawSearchAllByURIOptionImpl {
	res := &rawSearchAllByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRawSearchAllByURIOptions(opts ...RawSearchAllByURIOption) RawSearchAllByURIOptions {
	return makeRawSearchAllByURIOptionImpl(opts...)
}
