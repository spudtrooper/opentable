// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SearchAndSaveOption func(*searchAndSaveOptionImpl)

type SearchAndSaveOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchAndSaveVerbose(verbose bool) SearchAndSaveOption {
	return func(opts *searchAndSaveOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchAndSaveVerboseFlag(verbose *bool) SearchAndSaveOption {
	return func(opts *searchAndSaveOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
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
		opt(res)
	}
	return res
}

func MakeSearchAndSaveOptions(opts ...SearchAndSaveOption) SearchAndSaveOptions {
	return makeSearchAndSaveOptionImpl(opts...)
}
