// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package apihandler

type NewHandlerOption func(*newHandlerOptionImpl)

type NewHandlerOptions interface {
	RequiredFlagNames() []string
	HasRequiredFlagNames() bool
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

type newHandlerOptionImpl struct {
	requiredFlagNames     []string
	has_requiredFlagNames bool
}

func (n *newHandlerOptionImpl) RequiredFlagNames() []string { return n.requiredFlagNames }
func (n *newHandlerOptionImpl) HasRequiredFlagNames() bool  { return n.has_requiredFlagNames }

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
