// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type IncIntOption func(*incIntOptionImpl)

type IncIntOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func IncIntVerbose(verbose bool) IncIntOption {
	return func(opts *incIntOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func IncIntVerboseFlag(verbose *bool) IncIntOption {
	return func(opts *incIntOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type incIntOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (i *incIntOptionImpl) Verbose() bool    { return i.verbose }
func (i *incIntOptionImpl) HasVerbose() bool { return i.has_verbose }

func makeIncIntOptionImpl(opts ...IncIntOption) *incIntOptionImpl {
	res := &incIntOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeIncIntOptions(opts ...IncIntOption) IncIntOptions {
	return makeIncIntOptionImpl(opts...)
}
