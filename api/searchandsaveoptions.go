// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SearchAndSaveOption struct {
	f func(*searchAndSaveOptionImpl)
	s string
}

func (o SearchAndSaveOption) String() string { return o.s }

type SearchAndSaveOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchAndSaveVerbose(verbose bool) SearchAndSaveOption {
	return SearchAndSaveOption{func(opts *searchAndSaveOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchAndSaveVerbose(bool %+v)}", verbose)}
}
func SearchAndSaveVerboseFlag(verbose *bool) SearchAndSaveOption {
	return SearchAndSaveOption{func(opts *searchAndSaveOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchAndSaveVerbose(bool %+v)}", verbose)}
}

type searchAndSaveOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *searchAndSaveOptionImpl) Verbose() bool    { return s.verbose }
func (s *searchAndSaveOptionImpl) HasVerbose() bool { return s.has_verbose }

type SearchAndSaveParams struct {
	Term    string `json:"term" required:"true"`
	Verbose bool   `json:"verbose"`
}

func (o SearchAndSaveParams) Options() []SearchAndSaveOption {
	return []SearchAndSaveOption{
		SearchAndSaveVerbose(o.Verbose),
	}
}

func makeSearchAndSaveOptionImpl(opts ...SearchAndSaveOption) *searchAndSaveOptionImpl {
	res := &searchAndSaveOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchAndSaveOptions(opts ...SearchAndSaveOption) SearchAndSaveOptions {
	return makeSearchAndSaveOptionImpl(opts...)
}
