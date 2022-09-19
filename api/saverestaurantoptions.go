// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SaveRestaurantOption func(*saveRestaurantOptionImpl)

type SaveRestaurantOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SaveRestaurantVerbose(verbose bool) SaveRestaurantOption {
	return func(opts *saveRestaurantOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SaveRestaurantVerboseFlag(verbose *bool) SaveRestaurantOption {
	return func(opts *saveRestaurantOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type saveRestaurantOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *saveRestaurantOptionImpl) Verbose() bool    { return s.verbose }
func (s *saveRestaurantOptionImpl) HasVerbose() bool { return s.has_verbose }

func makeSaveRestaurantOptionImpl(opts ...SaveRestaurantOption) *saveRestaurantOptionImpl {
	res := &saveRestaurantOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSaveRestaurantOptions(opts ...SaveRestaurantOption) SaveRestaurantOptions {
	return makeSaveRestaurantOptionImpl(opts...)
}
