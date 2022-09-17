// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type HeaderUserProfileOption func(*headerUserProfileOptionImpl)

type HeaderUserProfileOptions interface {
	Tld() string
	HasTld() bool
	Verbose() bool
	HasVerbose() bool
}

func HeaderUserProfileTld(tld string) HeaderUserProfileOption {
	return func(opts *headerUserProfileOptionImpl) {
		opts.has_tld = true
		opts.tld = tld
	}
}
func HeaderUserProfileTldFlag(tld *string) HeaderUserProfileOption {
	return func(opts *headerUserProfileOptionImpl) {
		if tld == nil {
			return
		}
		opts.has_tld = true
		opts.tld = *tld
	}
}

func HeaderUserProfileVerbose(verbose bool) HeaderUserProfileOption {
	return func(opts *headerUserProfileOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func HeaderUserProfileVerboseFlag(verbose *bool) HeaderUserProfileOption {
	return func(opts *headerUserProfileOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
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

func makeHeaderUserProfileOptionImpl(opts ...HeaderUserProfileOption) *headerUserProfileOptionImpl {
	res := &headerUserProfileOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeHeaderUserProfileOptions(opts ...HeaderUserProfileOption) HeaderUserProfileOptions {
	return makeHeaderUserProfileOptionImpl(opts...)
}
