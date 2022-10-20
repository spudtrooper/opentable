// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SaveRawRestaurantDetailsFromIDOption struct {
	f func(*saveRawRestaurantDetailsFromIDOptionImpl)
	s string
}

func (o SaveRawRestaurantDetailsFromIDOption) String() string { return o.s }

type SaveRawRestaurantDetailsFromIDOptions interface {
	DebugFailures() bool
	HasDebugFailures() bool
	ReallyVerbose() bool
	HasReallyVerbose() bool
	Verbose() bool
	HasVerbose() bool
	ToRawRestaurantDetailsFromIDOptions() []RawRestaurantDetailsFromIDOption
	ToSaveRestaurantOptions() []SaveRestaurantOption
}

func SaveRawRestaurantDetailsFromIDDebugFailures(debugFailures bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDDebugFailures(bool %+v)}", debugFailures)}
}
func SaveRawRestaurantDetailsFromIDDebugFailuresFlag(debugFailures *bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDDebugFailures(bool %+v)}", debugFailures)}
}

func SaveRawRestaurantDetailsFromIDReallyVerbose(reallyVerbose bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_reallyVerbose = true
		opts.reallyVerbose = reallyVerbose
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDReallyVerbose(bool %+v)}", reallyVerbose)}
}
func SaveRawRestaurantDetailsFromIDReallyVerboseFlag(reallyVerbose *bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if reallyVerbose == nil {
			return
		}
		opts.has_reallyVerbose = true
		opts.reallyVerbose = *reallyVerbose
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDReallyVerbose(bool %+v)}", reallyVerbose)}
}

func SaveRawRestaurantDetailsFromIDVerbose(verbose bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDVerbose(bool %+v)}", verbose)}
}
func SaveRawRestaurantDetailsFromIDVerboseFlag(verbose *bool) SaveRawRestaurantDetailsFromIDOption {
	return SaveRawRestaurantDetailsFromIDOption{func(opts *saveRawRestaurantDetailsFromIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SaveRawRestaurantDetailsFromIDVerbose(bool %+v)}", verbose)}
}

type saveRawRestaurantDetailsFromIDOptionImpl struct {
	debugFailures     bool
	has_debugFailures bool
	reallyVerbose     bool
	has_reallyVerbose bool
	verbose           bool
	has_verbose       bool
}

func (s *saveRawRestaurantDetailsFromIDOptionImpl) DebugFailures() bool { return s.debugFailures }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasDebugFailures() bool {
	return s.has_debugFailures
}
func (s *saveRawRestaurantDetailsFromIDOptionImpl) ReallyVerbose() bool { return s.reallyVerbose }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasReallyVerbose() bool {
	return s.has_reallyVerbose
}
func (s *saveRawRestaurantDetailsFromIDOptionImpl) Verbose() bool    { return s.verbose }
func (s *saveRawRestaurantDetailsFromIDOptionImpl) HasVerbose() bool { return s.has_verbose }

type SaveRawRestaurantDetailsFromIDParams struct {
	DebugFailures bool   `json:"debug_failures"`
	ReallyVerbose bool   `json:"really_verbose"`
	RestID        string `json:"rest_id" required:"true"`
	Verbose       bool   `json:"verbose"`
}

func (o SaveRawRestaurantDetailsFromIDParams) Options() []SaveRawRestaurantDetailsFromIDOption {
	return []SaveRawRestaurantDetailsFromIDOption{
		SaveRawRestaurantDetailsFromIDDebugFailures(o.DebugFailures),
		SaveRawRestaurantDetailsFromIDReallyVerbose(o.ReallyVerbose),
		SaveRawRestaurantDetailsFromIDVerbose(o.Verbose),
	}
}

// ToRawRestaurantDetailsFromIDOptions converts SaveRawRestaurantDetailsFromIDOption to an array of RawRestaurantDetailsFromIDOption
func (o *saveRawRestaurantDetailsFromIDOptionImpl) ToRawRestaurantDetailsFromIDOptions() []RawRestaurantDetailsFromIDOption {
	return []RawRestaurantDetailsFromIDOption{
		RawRestaurantDetailsFromIDVerbose(o.Verbose()),
		RawRestaurantDetailsFromIDDebugFailures(o.DebugFailures()),
	}
}

// ToSaveRestaurantOptions converts SaveRawRestaurantDetailsFromIDOption to an array of SaveRestaurantOption
func (o *saveRawRestaurantDetailsFromIDOptionImpl) ToSaveRestaurantOptions() []SaveRestaurantOption {
	return []SaveRestaurantOption{
		SaveRestaurantVerbose(o.Verbose()),
		SaveRestaurantReallyVerbose(o.ReallyVerbose()),
	}
}

func makeSaveRawRestaurantDetailsFromIDOptionImpl(opts ...SaveRawRestaurantDetailsFromIDOption) *saveRawRestaurantDetailsFromIDOptionImpl {
	res := &saveRawRestaurantDetailsFromIDOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSaveRawRestaurantDetailsFromIDOptions(opts ...SaveRawRestaurantDetailsFromIDOption) SaveRawRestaurantDetailsFromIDOptions {
	return makeSaveRawRestaurantDetailsFromIDOptionImpl(opts...)
}
