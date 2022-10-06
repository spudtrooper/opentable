// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SaveRawRestaurantDetailsFromIDOption func(*saveRawRestaurantDetailsFromIDOptionImpl)

type SaveRawRestaurantDetailsFromIDOptions interface {
	Verbose() bool
	HasVerbose() bool
	ReallyVerbose() bool
	HasReallyVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
	ToSaveRestaurantOptions() []SaveRestaurantOption
	ToRawRestaurantDetailsFromIDOptions() []RawRestaurantDetailsFromIDOption
}

func SaveRawRestaurantDetailsFromIDVerbose(verbose bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SaveRawRestaurantDetailsFromIDVerboseFlag(verbose *bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func SaveRawRestaurantDetailsFromIDReallyVerbose(reallyVerbose bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_reallyVerbose = true
		opts.reallyVerbose = reallyVerbose
	}
}
func SaveRawRestaurantDetailsFromIDReallyVerboseFlag(reallyVerbose *bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if reallyVerbose == nil {
			return
		}
		opts.has_reallyVerbose = true
		opts.reallyVerbose = *reallyVerbose
	}
}

func SaveRawRestaurantDetailsFromIDDebugFailures(debugFailures bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func SaveRawRestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) SaveRawRestaurantDetailsFromIDOption {
	return func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

type saveRawRestaurantDetailsFromIDOptionImpl struct {
	verbose           bool
	has_verbose       bool
	reallyVerbose     bool
	has_reallyVerbose bool
	debugFailures     bool
	has_debugFailures bool
}

func (s *saveRawRestaurantDetailsFromIDOptionImpl) Verbose() bool       { return s.verbose }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasVerbose() bool    { return s.has_verbose }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) ReallyVerbose() bool { return s.reallyVerbose }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasReallyVerbose() bool {
	return s.has_reallyVerbose
}
func (s *saveRawRestaurantDetailsFromIDOptionImpl) DebugFailures() bool { return s.debugFailures }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasDebugFailures() bool {
	return s.has_debugFailures
}

type SaveRawRestaurantDetailsFromIDParams struct {
	RestID        string `json:"rest_id" required:"true"`
	Verbose       bool   `json:"verbose"`
	ReallyVerbose bool   `json:"really_verbose"`
	DebugFailures bool   `json:"debug_failures"`
}

func (o SaveRawRestaurantDetailsFromIDParams) Options() []SaveRawRestaurantDetailsFromIDOption {
	return []SaveRawRestaurantDetailsFromIDOption{
		SaveRawRestaurantDetailsFromIDVerbose(o.Verbose),
		SaveRawRestaurantDetailsFromIDReallyVerbose(o.ReallyVerbose),
		SaveRawRestaurantDetailsFromIDDebugFailures(o.DebugFailures),
	}
}

// ToSaveRestaurantOptions converts SaveRawRestaurantDetailsFromIDOption to an array of SaveRestaurantOption
func (o *saveRawRestaurantDetailsFromIDOptionImpl) ToSaveRestaurantOptions() []SaveRestaurantOption {
	return []SaveRestaurantOption{
		SaveRestaurantVerbose(o.Verbose()),
		SaveRestaurantReallyVerbose(o.ReallyVerbose()),
	}
}

// ToRawRestaurantDetailsFromIDOptions converts SaveRawRestaurantDetailsFromIDOption to an array of RawRestaurantDetailsFromIDOption
func (o *saveRawRestaurantDetailsFromIDOptionImpl) ToRawRestaurantDetailsFromIDOptions() []RawRestaurantDetailsFromIDOption {
	return []RawRestaurantDetailsFromIDOption{
		RawRestaurantDetailsFromIDVerbose(o.Verbose()),
		RawRestaurantDetailsFromIDDebugFailures(o.DebugFailures()),
	}
}

func makeSaveRawRestaurantDetailsFromIDOptionImpl(opts ...SaveRawRestaurantDetailsFromIDOption) *saveRawRestaurantDetailsFromIDOptionImpl {
	res := &saveRawRestaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSaveRawRestaurantDetailsFromIDOptions(opts ...SaveRawRestaurantDetailsFromIDOption) SaveRawRestaurantDetailsFromIDOptions {
	return makeSaveRawRestaurantDetailsFromIDOptionImpl(opts...)
}
