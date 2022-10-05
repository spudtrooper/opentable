// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RestaurantDetailsFromIDOption func(*restaurantDetailsFromIDOptionImpl)

type RestaurantDetailsFromIDOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
}

func RestaurantDetailsFromIDVerbose(verbose bool) RestaurantDetailsFromIDOption {
	return func(opts *restaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RestaurantDetailsFromIDVerboseFlag(verbose *bool) RestaurantDetailsFromIDOption {
	return func(opts *restaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RestaurantDetailsFromIDDebugFailures(debugFailures bool) RestaurantDetailsFromIDOption {
	return func(opts *restaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func RestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) RestaurantDetailsFromIDOption {
	return func(opts *restaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

type restaurantDetailsFromIDOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *restaurantDetailsFromIDOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsFromIDOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *restaurantDetailsFromIDOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsFromIDOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }

type RestaurantDetailsFromIDParams struct {
	Id            string `json:"id" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o RestaurantDetailsFromIDParams) Options() []RestaurantDetailsFromIDOption {
	return []RestaurantDetailsFromIDOption{
		RestaurantDetailsFromIDVerbose(o.Verbose),
		RestaurantDetailsFromIDDebugFailures(o.DebugFailures),
	}
}

func makeRestaurantDetailsFromIDOptionImpl(opts ...RestaurantDetailsFromIDOption) *restaurantDetailsFromIDOptionImpl {
	res := &restaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRestaurantDetailsFromIDOptions(opts ...RestaurantDetailsFromIDOption) RestaurantDetailsFromIDOptions {
	return makeRestaurantDetailsFromIDOptionImpl(opts...)
}
