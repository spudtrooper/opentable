// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type AddRestaurantsToSearchByURIsOption struct {
	f func(*addRestaurantsToSearchByURIsOptionImpl)
	s string
}

func (o AddRestaurantsToSearchByURIsOption) String() string { return o.s }

type AddRestaurantsToSearchByURIsOptions interface {
	Threads() int
	HasThreads() bool
	Verbose() bool
	HasVerbose() bool
}

func AddRestaurantsToSearchByURIsThreads(threads int) AddRestaurantsToSearchByURIsOption {
	return AddRestaurantsToSearchByURIsOption{func(opts *addRestaurantsToSearchByURIsOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.AddRestaurantsToSearchByURIsThreads(int %+v)", threads)}
}
func AddRestaurantsToSearchByURIsThreadsFlag(threads *int) AddRestaurantsToSearchByURIsOption {
	return AddRestaurantsToSearchByURIsOption{func(opts *addRestaurantsToSearchByURIsOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.AddRestaurantsToSearchByURIsThreads(int %+v)", threads)}
}

func AddRestaurantsToSearchByURIsVerbose(verbose bool) AddRestaurantsToSearchByURIsOption {
	return AddRestaurantsToSearchByURIsOption{func(opts *addRestaurantsToSearchByURIsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.AddRestaurantsToSearchByURIsVerbose(bool %+v)", verbose)}
}
func AddRestaurantsToSearchByURIsVerboseFlag(verbose *bool) AddRestaurantsToSearchByURIsOption {
	return AddRestaurantsToSearchByURIsOption{func(opts *addRestaurantsToSearchByURIsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.AddRestaurantsToSearchByURIsVerbose(bool %+v)", verbose)}
}

type addRestaurantsToSearchByURIsOptionImpl struct {
	threads     int
	has_threads bool
	verbose     bool
	has_verbose bool
}

func (a *addRestaurantsToSearchByURIsOptionImpl) Threads() int     { return a.threads }
func (a *addRestaurantsToSearchByURIsOptionImpl) HasThreads() bool { return a.has_threads }
func (a *addRestaurantsToSearchByURIsOptionImpl) Verbose() bool    { return a.verbose }
func (a *addRestaurantsToSearchByURIsOptionImpl) HasVerbose() bool { return a.has_verbose }

type AddRestaurantsToSearchByURIsParams struct {
	Threads int  `json:"threads"`
	Verbose bool `json:"verbose"`
}

func (o AddRestaurantsToSearchByURIsParams) Options() []AddRestaurantsToSearchByURIsOption {
	return []AddRestaurantsToSearchByURIsOption{
		AddRestaurantsToSearchByURIsThreads(o.Threads),
		AddRestaurantsToSearchByURIsVerbose(o.Verbose),
	}
}

func makeAddRestaurantsToSearchByURIsOptionImpl(opts ...AddRestaurantsToSearchByURIsOption) *addRestaurantsToSearchByURIsOptionImpl {
	res := &addRestaurantsToSearchByURIsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAddRestaurantsToSearchByURIsOptions(opts ...AddRestaurantsToSearchByURIsOption) AddRestaurantsToSearchByURIsOptions {
	return makeAddRestaurantsToSearchByURIsOptionImpl(opts...)
}
