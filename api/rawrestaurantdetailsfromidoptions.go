// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RawRestaurantDetailsFromIDOption func(*rawRestaurantDetailsFromIDOptionImpl)

type RawRestaurantDetailsFromIDOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
}

func RawRestaurantDetailsFromIDVerbose(verbose bool) RawRestaurantDetailsFromIDOption {
	return func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RawRestaurantDetailsFromIDVerboseFlag(verbose *bool) RawRestaurantDetailsFromIDOption {
	return func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RawRestaurantDetailsFromIDDebugFailures(debugFailures bool) RawRestaurantDetailsFromIDOption {
	return func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func RawRestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) RawRestaurantDetailsFromIDOption {
	return func(opts *rawRestaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

type rawRestaurantDetailsFromIDOptionImpl struct {
	verbose           bool
	has_verbose       bool
	debugFailures     bool
	has_debugFailures bool
}

func (r *rawRestaurantDetailsFromIDOptionImpl) Verbose() bool          { return r.verbose }
func (r *rawRestaurantDetailsFromIDOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *rawRestaurantDetailsFromIDOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *rawRestaurantDetailsFromIDOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }

type RawRestaurantDetailsFromIDParams struct {
	Id            string `json:"id" required:"true"`
	Verbose       bool   `json:"verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o RawRestaurantDetailsFromIDParams) Options() []RawRestaurantDetailsFromIDOption {
	return []RawRestaurantDetailsFromIDOption{
		RawRestaurantDetailsFromIDVerbose(o.Verbose),
		RawRestaurantDetailsFromIDDebugFailures(o.DebugFailures),
	}
}

func makeRawRestaurantDetailsFromIDOptionImpl(opts ...RawRestaurantDetailsFromIDOption) *rawRestaurantDetailsFromIDOptionImpl {
	res := &rawRestaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRawRestaurantDetailsFromIDOptions(opts ...RawRestaurantDetailsFromIDOption) RawRestaurantDetailsFromIDOptions {
	return makeRawRestaurantDetailsFromIDOptionImpl(opts...)
}
