// Package apihandler generates a CLI and frontend API server to access a third party API.
package apihandler

type HandlerFn func(ctx EvalContext) (interface{}, error)

type Handler interface{}

type handler struct {
	name              string
	requiredFlagNames []string
	fn                HandlerFn
}

//go:generate genopts --function NewHandler "requiredFlagNames:[]string"
func NewHandler(name string, fn HandlerFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	return &handler{
		name:              name,
		requiredFlagNames: opts.RequiredFlagNames(),
		fn:                fn,
	}
}
