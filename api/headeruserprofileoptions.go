// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type HeaderUserProfileOption struct {
	f func(*headerUserProfileOptionImpl)
	s string
}

func (o HeaderUserProfileOption) String() string { return o.s }

type HeaderUserProfileOptions interface {
	Tld() string
	HasTld() bool
	Verbose() bool
	HasVerbose() bool
}

func HeaderUserProfileTld(tld string) HeaderUserProfileOption {
	return HeaderUserProfileOption{func(opts *headerUserProfileOptionImpl) {
		opts.has_tld = true
		opts.tld = tld
	}, fmt.Sprintf("api.HeaderUserProfileTld(string %+v)}", tld)}
}
func HeaderUserProfileTldFlag(tld *string) HeaderUserProfileOption {
	return HeaderUserProfileOption{func(opts *headerUserProfileOptionImpl) {
		if tld == nil {
			return
		}
		opts.has_tld = true
		opts.tld = *tld
	}, fmt.Sprintf("api.HeaderUserProfileTld(string %+v)}", tld)}
}

func HeaderUserProfileVerbose(verbose bool) HeaderUserProfileOption {
	return HeaderUserProfileOption{func(opts *headerUserProfileOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.HeaderUserProfileVerbose(bool %+v)}", verbose)}
}
func HeaderUserProfileVerboseFlag(verbose *bool) HeaderUserProfileOption {
	return HeaderUserProfileOption{func(opts *headerUserProfileOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.HeaderUserProfileVerbose(bool %+v)}", verbose)}
}

type headerUserProfileOptionImpl struct {
	tld         string
	has_tld     bool
	verbose     bool
	has_verbose bool
}

func (h *headerUserProfileOptionImpl) Tld() string      { return h.tld }
func (h *headerUserProfileOptionImpl) HasTld() bool     { return h.has_tld }
func (h *headerUserProfileOptionImpl) Verbose() bool    { return h.verbose }
func (h *headerUserProfileOptionImpl) HasVerbose() bool { return h.has_verbose }

type HeaderUserProfileParams struct {
	Tld     string `json:"tld"`
	Verbose bool   `json:"verbose"`
}

func (o HeaderUserProfileParams) Options() []HeaderUserProfileOption {
	return []HeaderUserProfileOption{
		HeaderUserProfileTld(o.Tld),
		HeaderUserProfileVerbose(o.Verbose),
	}
}

func makeHeaderUserProfileOptionImpl(opts ...HeaderUserProfileOption) *headerUserProfileOptionImpl {
	res := &headerUserProfileOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeHeaderUserProfileOptions(opts ...HeaderUserProfileOption) HeaderUserProfileOptions {
	return makeHeaderUserProfileOptionImpl(opts...)
}
