// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SaveRestaurantOption struct {
	f func(*saveRestaurantOptionImpl)
	s string
}

func (o SaveRestaurantOption) String() string { return o.s }

type SaveRestaurantOptions interface {
	ReallyVerbose() bool
	HasReallyVerbose() bool
	Verbose() bool
	HasVerbose() bool
}

func SaveRestaurantReallyVerbose(reallyVerbose bool) SaveRestaurantOption {
	return SaveRestaurantOption{func(opts *saveRestaurantOptionImpl) {
		opts.has_reallyVerbose = true
		opts.reallyVerbose = reallyVerbose
	}, fmt.Sprintf("api.SaveRestaurantReallyVerbose(bool %+v)", reallyVerbose)}
}
func SaveRestaurantReallyVerboseFlag(reallyVerbose *bool) SaveRestaurantOption {
	return SaveRestaurantOption{func(opts *saveRestaurantOptionImpl) {
		if reallyVerbose == nil {
			return
		}
		opts.has_reallyVerbose = true
		opts.reallyVerbose = *reallyVerbose
	}, fmt.Sprintf("api.SaveRestaurantReallyVerbose(bool %+v)", reallyVerbose)}
}

func SaveRestaurantVerbose(verbose bool) SaveRestaurantOption {
	return SaveRestaurantOption{func(opts *saveRestaurantOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SaveRestaurantVerbose(bool %+v)", verbose)}
}
func SaveRestaurantVerboseFlag(verbose *bool) SaveRestaurantOption {
	return SaveRestaurantOption{func(opts *saveRestaurantOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SaveRestaurantVerbose(bool %+v)", verbose)}
}

type saveRestaurantOptionImpl struct {
	reallyVerbose     bool
	has_reallyVerbose bool
	verbose           bool
	has_verbose       bool
}

func (s *saveRestaurantOptionImpl) ReallyVerbose() bool    { return s.reallyVerbose }
func (s *saveRestaurantOptionImpl) HasReallyVerbose() bool { return s.has_reallyVerbose }
func (s *saveRestaurantOptionImpl) Verbose() bool          { return s.verbose }
func (s *saveRestaurantOptionImpl) HasVerbose() bool       { return s.has_verbose }

func makeSaveRestaurantOptionImpl(opts ...SaveRestaurantOption) *saveRestaurantOptionImpl {
	res := &saveRestaurantOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSaveRestaurantOptions(opts ...SaveRestaurantOption) SaveRestaurantOptions {
	return makeSaveRestaurantOptionImpl(opts...)
}
