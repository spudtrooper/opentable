// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SearchRestaurantFromQueueOption struct {
	f func(*searchRestaurantFromQueueOptionImpl)
	s string
}

func (o SearchRestaurantFromQueueOption) String() string { return o.s }

type SearchRestaurantFromQueueOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchRestaurantFromQueueVerbose(verbose bool) SearchRestaurantFromQueueOption {
	return SearchRestaurantFromQueueOption{func(opts *searchRestaurantFromQueueOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchRestaurantFromQueueVerbose(bool %+v)", verbose)}
}
func SearchRestaurantFromQueueVerboseFlag(verbose *bool) SearchRestaurantFromQueueOption {
	return SearchRestaurantFromQueueOption{func(opts *searchRestaurantFromQueueOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchRestaurantFromQueueVerbose(bool %+v)", verbose)}
}

type searchRestaurantFromQueueOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *searchRestaurantFromQueueOptionImpl) Verbose() bool    { return s.verbose }
func (s *searchRestaurantFromQueueOptionImpl) HasVerbose() bool { return s.has_verbose }

type SearchRestaurantFromQueueParams struct {
	Verbose bool `json:"verbose"`
}

func (o SearchRestaurantFromQueueParams) Options() []SearchRestaurantFromQueueOption {
	return []SearchRestaurantFromQueueOption{
		SearchRestaurantFromQueueVerbose(o.Verbose),
	}
}

func makeSearchRestaurantFromQueueOptionImpl(opts ...SearchRestaurantFromQueueOption) *searchRestaurantFromQueueOptionImpl {
	res := &searchRestaurantFromQueueOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchRestaurantFromQueueOptions(opts ...SearchRestaurantFromQueueOption) SearchRestaurantFromQueueOptions {
	return makeSearchRestaurantFromQueueOptionImpl(opts...)
}
