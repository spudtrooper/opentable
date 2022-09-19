// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type SearchByURIAndSaveOption func(*searchByURIAndSaveOptionImpl)

type SearchByURIAndSaveOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func SearchByURIAndSaveVerbose(verbose bool) SearchByURIAndSaveOption {
	return func(opts *searchByURIAndSaveOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchByURIAndSaveVerboseFlag(verbose *bool) SearchByURIAndSaveOption {
	return func(opts *searchByURIAndSaveOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type searchByURIAndSaveOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (s *searchByURIAndSaveOptionImpl) Verbose() bool    { return s.verbose }
func (s *searchByURIAndSaveOptionImpl) HasVerbose() bool { return s.has_verbose }

func makeSearchByURIAndSaveOptionImpl(opts ...SearchByURIAndSaveOption) *searchByURIAndSaveOptionImpl {
	res := &searchByURIAndSaveOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchByURIAndSaveOptions(opts ...SearchByURIAndSaveOption) SearchByURIAndSaveOptions {
	return makeSearchByURIAndSaveOptionImpl(opts...)
}
