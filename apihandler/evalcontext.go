package apihandler

import (
	"context"
	"log"
	"net/http"
)

type EvalContext interface {
	Context() context.Context
	String(name string) (string, bool)
	MustString(name string) (string, bool)
	Bool(name string) bool
}

type cliEvalContext struct {
	ctx         context.Context
	stringFlags map[string]*string
	boolFlags   map[string]*bool
}

func (c *cliEvalContext) Context() context.Context { return c.ctx }

func (c *cliEvalContext) String(name string) (string, bool) {
	flag, ok := c.stringFlags[name]
	if !ok {
		return "", false
	}
	return *flag, true
}

func (c *cliEvalContext) MustString(name string) (string, bool) {
	val, ok := c.String(name)
	if !ok || val == "" {
		log.Fatalf("--%s required", name)
	}
	return val, true
}

func (c *cliEvalContext) Bool(name string) bool {
	flag, ok := c.boolFlags[name]
	if !ok {
		return false
	}
	return *flag
}

type serverEvalContext struct {
	ctx context.Context
	w   http.ResponseWriter
	req *http.Request
}

func (c *serverEvalContext) Context() context.Context { return c.ctx }

func (c *serverEvalContext) String(name string) (string, bool) {
	val := getStringURLParam(c.req, name)
	ok := val != "" //TODO: ???
	return val, ok
}

func (c *serverEvalContext) MustString(name string) (string, bool) {
	return getStringURLParamOrDie(c.w, c.req, name)
}

func (c *serverEvalContext) Bool(name string) bool {
	return getBoolURLParam(c.req, name)
}
