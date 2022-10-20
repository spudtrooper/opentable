// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type RawListAllByURIOption struct {
	f func(*rawListAllByURIOptionImpl)
	s string
}

func (o RawListAllByURIOption) String() string { return o.s }

type RawListAllByURIOptions interface {
	StartPage() int
	HasStartPage() bool
	Sync() bool
	HasSync() bool
	Threads() int
	HasThreads() bool
	Verbose() bool
	HasVerbose() bool
}

func RawListAllByURIStartPage(startPage int) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}, fmt.Sprintf("api.RawListAllByURIStartPage(int %+v)}", startPage)}
}
func RawListAllByURIStartPageFlag(startPage *int) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}, fmt.Sprintf("api.RawListAllByURIStartPage(int %+v)}", startPage)}
}

func RawListAllByURISync(sync bool) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		opts.has_sync = true
		opts.sync = sync
	}, fmt.Sprintf("api.RawListAllByURISync(bool %+v)}", sync)}
}
func RawListAllByURISyncFlag(sync *bool) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		if sync == nil {
			return
		}
		opts.has_sync = true
		opts.sync = *sync
	}, fmt.Sprintf("api.RawListAllByURISync(bool %+v)}", sync)}
}

func RawListAllByURIThreads(threads int) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.RawListAllByURIThreads(int %+v)}", threads)}
}
func RawListAllByURIThreadsFlag(threads *int) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.RawListAllByURIThreads(int %+v)}", threads)}
}

func RawListAllByURIVerbose(verbose bool) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RawListAllByURIVerbose(bool %+v)}", verbose)}
}
func RawListAllByURIVerboseFlag(verbose *bool) RawListAllByURIOption {
	return RawListAllByURIOption{func(opts *rawListAllByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RawListAllByURIVerbose(bool %+v)}", verbose)}
}

type rawListAllByURIOptionImpl struct {
	verbose       bool
	has_verbose   bool
	startPage     int
	has_startPage bool
	threads       int
	has_threads   bool
	sync          bool
	has_sync      bool
}

func (r *rawListAllByURIOptionImpl) StartPage() int     { return r.startPage }
func (r *rawListAllByURIOptionImpl) HasStartPage() bool { return r.has_startPage }
func (r *rawListAllByURIOptionImpl) Sync() bool         { return r.sync }
func (r *rawListAllByURIOptionImpl) HasSync() bool      { return r.has_sync }
func (r *rawListAllByURIOptionImpl) Threads() int       { return r.threads }
func (r *rawListAllByURIOptionImpl) HasThreads() bool   { return r.has_threads }
func (r *rawListAllByURIOptionImpl) Verbose() bool      { return r.verbose }
func (r *rawListAllByURIOptionImpl) HasVerbose() bool   { return r.has_verbose }

type RawListAllByURIParams struct {
	StartPage int    `json:"start_page"`
	Sync      bool   `json:"sync"`
	Threads   int    `json:"threads"`
	Uri       string `json:"uri" required:"true"`
	Verbose   bool   `json:"verbose"`
}

func (o RawListAllByURIParams) Options() []RawListAllByURIOption {
	return []RawListAllByURIOption{
		RawListAllByURIStartPage(o.StartPage),
		RawListAllByURISync(o.Sync),
		RawListAllByURIThreads(o.Threads),
		RawListAllByURIVerbose(o.Verbose),
	}
}

func makeRawListAllByURIOptionImpl(opts ...RawListAllByURIOption) *rawListAllByURIOptionImpl {
	res := &rawListAllByURIOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRawListAllByURIOptions(opts ...RawListAllByURIOption) RawListAllByURIOptions {
	return makeRawListAllByURIOptionImpl(opts...)
}
