// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package apihandler

type NewHandlerOption func(*newHandlerOptionImpl)

type NewHandlerOptions interface {
	RequiredFlagNames() []string
	HasRequiredFlagNames() bool
	CliOnly() bool
	HasCliOnly() bool
}

func NewHandlerRequiredFlagNames(requiredFlagNames []string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_requiredFlagNames = true
		opts.requiredFlagNames = requiredFlagNames
	}
}
func NewHandlerRequiredFlagNamesFlag(requiredFlagNames *[]string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if requiredFlagNames == nil {
			return
		}
		opts.has_requiredFlagNames = true
		opts.requiredFlagNames = *requiredFlagNames
	}
}

func NewHandlerCliOnly(cliOnly bool) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_cliOnly = true
		opts.cliOnly = cliOnly
	}
}
func NewHandlerCliOnlyFlag(cliOnly *bool) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if cliOnly == nil {
			return
		}
		opts.has_cliOnly = true
		opts.cliOnly = *cliOnly
	}
}

type newHandlerOptionImpl struct {
	requiredFlagNames     []string
	has_requiredFlagNames bool
	cliOnly               bool
	has_cliOnly           bool
}

func (n *newHandlerOptionImpl) RequiredFlagNames() []string { return n.requiredFlagNames }
func (n *newHandlerOptionImpl) HasRequiredFlagNames() bool  { return n.has_requiredFlagNames }
func (n *newHandlerOptionImpl) CliOnly() bool               { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool            { return n.has_cliOnly }

func makeNewHandlerOptionImpl(opts ...NewHandlerOption) *newHandlerOptionImpl {
	res := &newHandlerOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeNewHandlerOptions(opts ...NewHandlerOption) NewHandlerOptions {
	return makeNewHandlerOptionImpl(opts...)
}
