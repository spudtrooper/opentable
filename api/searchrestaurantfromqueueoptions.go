// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SearchRestaurantFromQueueOption func(*searchRestaurantFromQueueOptionImpl)

type SearchRestaurantFromQueueOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchRestaurantFromQueueVerbose(verbose bool) SearchRestaurantFromQueueOption {
	return func(opts *searchRestaurantFromQueueOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchRestaurantFromQueueVerboseFlag(verbose *bool) SearchRestaurantFromQueueOption {
	return func(opts *searchRestaurantFromQueueOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type searchRestaurantFromQueueOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *searchRestaurantFromQueueOptionImpl) Verbose() bool    { return s.verbose }
func (s *searchRestaurantFromQueueOptionImpl) HasVerbose() bool { return s.has_verbose }

func makeSearchRestaurantFromQueueOptionImpl(opts ...SearchRestaurantFromQueueOption) *searchRestaurantFromQueueOptionImpl {
	res := &searchRestaurantFromQueueOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchRestaurantFromQueueOptions(opts ...SearchRestaurantFromQueueOption) SearchRestaurantFromQueueOptions {
	return makeSearchRestaurantFromQueueOptionImpl(opts...)
}
