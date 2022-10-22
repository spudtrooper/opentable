// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type SearchByURIAndSaveOption struct {
	f func(*searchByURIAndSaveOptionImpl)
	s string
}

func (o SearchByURIAndSaveOption) String() string { return o.s }

type SearchByURIAndSaveOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchByURIAndSaveVerbose(verbose bool) SearchByURIAndSaveOption {
	return SearchByURIAndSaveOption{func(opts *searchByURIAndSaveOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchByURIAndSaveVerbose(bool %+v)", verbose)}
}
func SearchByURIAndSaveVerboseFlag(verbose *bool) SearchByURIAndSaveOption {
	return SearchByURIAndSaveOption{func(opts *searchByURIAndSaveOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchByURIAndSaveVerbose(bool %+v)", verbose)}
}

type searchByURIAndSaveOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *searchByURIAndSaveOptionImpl) Verbose() bool    { return s.verbose }
func (s *searchByURIAndSaveOptionImpl) HasVerbose() bool { return s.has_verbose }

type SearchByURIAndSaveParams struct {
	Uri     string `json:"uri" required:"true"`
	Verbose bool   `json:"verbose"`
}

func (o SearchByURIAndSaveParams) Options() []SearchByURIAndSaveOption {
	return []SearchByURIAndSaveOption{
		SearchByURIAndSaveVerbose(o.Verbose),
	}
}

func makeSearchByURIAndSaveOptionImpl(opts ...SearchByURIAndSaveOption) *searchByURIAndSaveOptionImpl {
	res := &searchByURIAndSaveOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchByURIAndSaveOptions(opts ...SearchByURIAndSaveOption) SearchByURIAndSaveOptions {
	return makeSearchByURIAndSaveOptionImpl(opts...)
}
