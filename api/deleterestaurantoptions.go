// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type DeleteRestaurantOption func(*deleteRestaurantOptionImpl)

type DeleteRestaurantOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func DeleteRestaurantVerbose(verbose bool) DeleteRestaurantOption {
	return func(opts *deleteRestaurantOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func DeleteRestaurantVerboseFlag(verbose *bool) DeleteRestaurantOption {
	return func(opts *deleteRestaurantOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type deleteRestaurantOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (d *deleteRestaurantOptionImpl) Verbose() bool    { return d.verbose }
func (d *deleteRestaurantOptionImpl) HasVerbose() bool { return d.has_verbose }

func makeDeleteRestaurantOptionImpl(opts ...DeleteRestaurantOption) *deleteRestaurantOptionImpl {
	res := &deleteRestaurantOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeDeleteRestaurantOptions(opts ...DeleteRestaurantOption) DeleteRestaurantOptions {
	return makeDeleteRestaurantOptionImpl(opts...)
}
