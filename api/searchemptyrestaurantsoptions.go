// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type SearchEmptyRestaurantsOption func(*searchEmptyRestaurantsOptionImpl)

type SearchEmptyRestaurantsOptions interface {
	Threads() int
	HasThreads() bool
	Verbose() bool
	HasVerbose() bool
	Sleep() time.Duration
	HasSleep() bool
}

func SearchEmptyRestaurantsThreads(threads int) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}
}
func SearchEmptyRestaurantsThreadsFlag(threads *int) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}
}

func SearchEmptyRestaurantsVerbose(verbose bool) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchEmptyRestaurantsVerboseFlag(verbose *bool) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func SearchEmptyRestaurantsSleep(sleep time.Duration) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		opts.has_sleep = true
		opts.sleep = sleep
	}
}
func SearchEmptyRestaurantsSleepFlag(sleep *time.Duration) SearchEmptyRestaurantsOption {
	return func(opts *searchEmptyRestaurantsOptionImpl) {
		if sleep == nil {
			return
		}
		opts.has_sleep = true
		opts.sleep = *sleep
	}
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
		opt(res)
	}
	return res
}

func MakeSearchEmptyRestaurantsOptions(opts ...SearchEmptyRestaurantsOption) SearchEmptyRestaurantsOptions {
	return makeSearchEmptyRestaurantsOptionImpl(opts...)
}
