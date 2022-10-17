// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type SearchEmptyRestaurantsOption struct {
	f func(*searchEmptyRestaurantsOptionImpl)
	s string
}

func (o SearchEmptyRestaurantsOption) String() string { return o.s }

type SearchEmptyRestaurantsOptions interface {
	Threads() int
	HasThreads() bool
	Verbose() bool
	HasVerbose() bool
	Sleep() time.Duration
	HasSleep() bool
}

func SearchEmptyRestaurantsThreads(threads int) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.SearchEmptyRestaurantsThreads(int %+v)}", threads)}
}
func SearchEmptyRestaurantsThreadsFlag(threads *int) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.SearchEmptyRestaurantsThreads(int %+v)}", threads)}
}

func SearchEmptyRestaurantsVerbose(verbose bool) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchEmptyRestaurantsVerbose(bool %+v)}", verbose)}
}
func SearchEmptyRestaurantsVerboseFlag(verbose *bool) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchEmptyRestaurantsVerbose(bool %+v)}", verbose)}
}

func SearchEmptyRestaurantsSleep(sleep time.Duration) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_sleep = true
		opts.sleep = sleep
	}, fmt.Sprintf("api.SearchEmptyRestaurantsSleep(time.Duration %+v)}", sleep)}
}
func SearchEmptyRestaurantsSleepFlag(sleep *time.Duration) SearchEmptyRestaurantsOption {
	return SearchEmptyRestaurantsOption{func(opts *searchEmptyRestaurantsOptionImpl) {
		if sleep == nil {
			return
		}
		opts.has_sleep = true
		opts.sleep = *sleep
	}, fmt.Sprintf("api.SearchEmptyRestaurantsSleep(time.Duration %+v)}", sleep)}
}

type searchEmptyRestaurantsOptionImpl struct {
	threads     int
	has_threads bool
	verbose     bool
	has_verbose bool
	sleep       time.Duration
	has_sleep   bool
}

func (s *searchEmptyRestaurantsOptionImpl) Threads() int         { return s.threads }
func (s *searchEmptyRestaurantsOptionImpl) HasThreads() bool     { return s.has_threads }
func (s *searchEmptyRestaurantsOptionImpl) Verbose() bool        { return s.verbose }
func (s *searchEmptyRestaurantsOptionImpl) HasVerbose() bool     { return s.has_verbose }
func (s *searchEmptyRestaurantsOptionImpl) Sleep() time.Duration { return s.sleep }
func (s *searchEmptyRestaurantsOptionImpl) HasSleep() bool       { return s.has_sleep }

type SearchEmptyRestaurantsParams struct {
	Threads int           `json:"threads"`
	Verbose bool          `json:"verbose"`
	Sleep   time.Duration `json:"sleep"`
}

func (o SearchEmptyRestaurantsParams) Options() []SearchEmptyRestaurantsOption {
	return []SearchEmptyRestaurantsOption{
		SearchEmptyRestaurantsThreads(o.Threads),
		SearchEmptyRestaurantsVerbose(o.Verbose),
		SearchEmptyRestaurantsSleep(o.Sleep),
	}
}

func makeSearchEmptyRestaurantsOptionImpl(opts ...SearchEmptyRestaurantsOption) *searchEmptyRestaurantsOptionImpl {
	res := &searchEmptyRestaurantsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchEmptyRestaurantsOptions(opts ...SearchEmptyRestaurantsOption) SearchEmptyRestaurantsOptions {
	return makeSearchEmptyRestaurantsOptionImpl(opts...)
}
