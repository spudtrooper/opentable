// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RestaurantDetailsByURIOption func(*restaurantDetailsByURIOptionImpl)

type RestaurantDetailsByURIOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
}

func RestaurantDetailsByURIVerbose(verbose bool) RestaurantDetailsByURIOption {
	return func(opts *restaurantDetailsByURIOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RestaurantDetailsByURIVerboseFlag(verbose *bool) RestaurantDetailsByURIOption {
	return func(opts *restaurantDetailsByURIOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RestaurantDetailsByURIDebugFailures(debugFailures bool) RestaurantDetailsByURIOption {
	return func(opts *restaurantDetailsByURIOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func RestaurantDetailsByURIDebugFailuresFlag(debugFailures *bool) RestaurantDetailsByURIOption {
	return func(opts *restaurantDetailsByURIOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

type restaurantDetailsByURIOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *restaurantDetailsByURIOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsByURIOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *restaurantDetailsByURIOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsByURIOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }

type RestaurantDetailsByURIParams struct {
	Uri           string `json:"uri" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o RestaurantDetailsByURIParams) Options() []RestaurantDetailsByURIOption {
	return []RestaurantDetailsByURIOption{
		RestaurantDetailsByURIVerbose(o.Verbose),
		RestaurantDetailsByURIDebugFailures(o.DebugFailures),
	}
}

func makeRestaurantDetailsByURIOptionImpl(opts ...RestaurantDetailsByURIOption) *restaurantDetailsByURIOptionImpl {
	res := &restaurantDetailsByURIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRestaurantDetailsByURIOptions(opts ...RestaurantDetailsByURIOption) RestaurantDetailsByURIOptions {
	return makeRestaurantDetailsByURIOptionImpl(opts...)
}
