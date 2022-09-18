// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RawListAllByURIOption func(*rawListAllByURIOptionImpl)

type RawListAllByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	StartPage() int
	HasStartPage() bool
	Threads() int
	HasThreads() bool
	Sync() bool
	HasSync() bool
}

func RawListAllByURIVerbose(verbose bool) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RawListAllByURIVerboseFlag(verbose *bool) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RawListAllByURIStartPage(startPage int) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}
}
func RawListAllByURIStartPageFlag(startPage *int) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}
}

func RawListAllByURIThreads(threads int) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}
}
func RawListAllByURIThreadsFlag(threads *int) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}
}

func RawListAllByURISync(sync bool) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		opts.has_sync = true
		opts.sync = sync
	}
}
func RawListAllByURISyncFlag(sync *bool) RawListAllByURIOption {
	return func(opts *rawListAllByURIOptionImpl) {
		if sync == nil {
			return
		}
		opts.has_sync = true
		opts.sync = *sync
	}
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

func (r *rawListAllByURIOptionImpl) Verbose() bool      { return r.verbose }
func (r *rawListAllByURIOptionImpl) HasVerbose() bool   { return r.has_verbose }
func (r *rawListAllByURIOptionImpl) StartPage() int     { return r.startPage }
func (r *rawListAllByURIOptionImpl) HasStartPage() bool { return r.has_startPage }
func (r *rawListAllByURIOptionImpl) Threads() int       { return r.threads }
func (r *rawListAllByURIOptionImpl) HasThreads() bool   { return r.has_threads }
func (r *rawListAllByURIOptionImpl) Sync() bool         { return r.sync }
func (r *rawListAllByURIOptionImpl) HasSync() bool      { return r.has_sync }

func makeRawListAllByURIOptionImpl(opts ...RawListAllByURIOption) *rawListAllByURIOptionImpl {
	res := &rawListAllByURIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRawListAllByURIOptions(opts ...RawListAllByURIOption) RawListAllByURIOptions {
	return makeRawListAllByURIOptionImpl(opts...)
}
