// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type FindMenuItemOption func(*findMenuItemOptionImpl)

type FindMenuItemOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func FindMenuItemVerbose(verbose bool) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func FindMenuItemVerboseFlag(verbose *bool) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

type findMenuItemOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (f *findMenuItemOptionImpl) Verbose() bool    { return f.verbose }
func (f *findMenuItemOptionImpl) HasVerbose() bool { return f.has_verbose }

func makeFindMenuItemOptionImpl(opts ...FindMenuItemOption) *findMenuItemOptionImpl {
	res := &findMenuItemOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeFindMenuItemOptions(opts ...FindMenuItemOption) FindMenuItemOptions {
	return makeFindMenuItemOptionImpl(opts...)
}
