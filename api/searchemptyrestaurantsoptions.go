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
	Sleep() time.Duration
	HasSleep() bool
	Threads() int
	HasThreads() bool
	Verbose() bool
	HasVerbose() bool
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

type searchEmptyRestaurantsOptionImpl struct {
	sleep       time.Duration
	has_sleep   bool
	threads     int
	has_threads bool
	verbose     bool
	has_verbose bool
}

func (s *searchEmptyRestaurantsOptionImpl) Sleep() time.Duration { return s.sleep }
func (s *searchEmptyRestaurantsOptionImpl) HasSleep() bool       { return s.has_sleep }
func (s *searchEmptyRestaurantsOptionImpl) Threads() int         { return s.threads }
func (s *searchEmptyRestaurantsOptionImpl) HasThreads() bool     { return s.has_threads }
func (s *searchEmptyRestaurantsOptionImpl) Verbose() bool        { return s.verbose }
func (s *searchEmptyRestaurantsOptionImpl) HasVerbose() bool     { return s.has_verbose }

type SearchEmptyRestaurantsParams struct {
	Sleep   time.Duration `json:"sleep"`
	Threads int           `json:"threads"`
	Verbose bool          `json:"verbose"`
}

func (o SearchEmptyRestaurantsParams) Options() []SearchEmptyRestaurantsOption {
	return []SearchEmptyRestaurantsOption{
		SearchEmptyRestaurantsSleep(o.Sleep),
		SearchEmptyRestaurantsThreads(o.Threads),
		SearchEmptyRestaurantsVerbose(o.Verbose),
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
