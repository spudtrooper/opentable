// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type LolzViewAllByURIOption func(*lolzViewAllByURIOptionImpl)

type LolzViewAllByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	StartPage() int
	HasStartPage() bool
	Threads() int
	HasThreads() bool
}

func LolzViewAllByURIVerbose(verbose bool) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func LolzViewAllByURIVerboseFlag(verbose *bool) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func LolzViewAllByURIStartPage(startPage int) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}
}
func LolzViewAllByURIStartPageFlag(startPage *int) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}
}

func LolzViewAllByURIThreads(threads int) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}
}
func LolzViewAllByURIThreadsFlag(threads *int) LolzViewAllByURIOption {
	return func(opts *lolzViewAllByURIOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}
}

type lolzViewAllByURIOptionImpl struct {
	verbose       bool
	has_verbose   bool
	startPage     int
	has_startPage bool
	threads       int
	has_threads   bool
}

func (l *lolzViewAllByURIOptionImpl) Verbose() bool      { return l.verbose }
func (l *lolzViewAllByURIOptionImpl) HasVerbose() bool   { return l.has_verbose }
func (l *lolzViewAllByURIOptionImpl) StartPage() int     { return l.startPage }
func (l *lolzViewAllByURIOptionImpl) HasStartPage() bool { return l.has_startPage }
func (l *lolzViewAllByURIOptionImpl) Threads() int       { return l.threads }
func (l *lolzViewAllByURIOptionImpl) HasThreads() bool   { return l.has_threads }

type LolzViewAllByURIParams struct {
	Uri       string `json:"uri" required:"true"`
	Verbose   bool   `json:"verbose"`
	StartPage int    `json:"start_page"`
	Threads   int    `json:"threads"`
}

func (o LolzViewAllByURIParams) Options() []LolzViewAllByURIOption {
	return []LolzViewAllByURIOption{
		LolzViewAllByURIVerbose(o.Verbose),
		LolzViewAllByURIStartPage(o.StartPage),
		LolzViewAllByURIThreads(o.Threads),
	}
}

func makeLolzViewAllByURIOptionImpl(opts ...LolzViewAllByURIOption) *lolzViewAllByURIOptionImpl {
	res := &lolzViewAllByURIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeLolzViewAllByURIOptions(opts ...LolzViewAllByURIOption) LolzViewAllByURIOptions {
	return makeLolzViewAllByURIOptionImpl(opts...)
}
