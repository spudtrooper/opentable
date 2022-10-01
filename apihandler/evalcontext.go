package apihandler

import (
	"context"
	"log"
	"net/http"
	"time"
)

type EvalContext interface {
	Context() context.Context
	String(name string) (string, bool)
	MustString(name string) (string, bool)
	Bool(name string) bool
	Int(name string) int
	Duration(name string) time.Duration
}

type cliEvalContext struct {
	ctx         context.Context
	stringFlags map[string]*string
	boolFlags   map[string]*bool
	intFlags    map[string]*int
	durFlags    map[string]*time.Duration
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

func (c *cliEvalContext) Int(name string) int {
	flag, ok := c.intFlags[name]
	if !ok {
		return 0
	}
	return *flag
}

func (c *cliEvalContext) Duration(name string) time.Duration {
	flag, ok := c.durFlags[name]
	if !ok {
		return 0
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

func (c *serverEvalContext) Bool(name string) bool { return getBoolURLParam(c.req, name) }
func (c *serverEvalContext) Int(name string) int   { return getIntURLParam(c.req, name) }
func (c *serverEvalContext) Duration(name string) time.Duration {
	return time.Duration(getIntURLParam(c.req, name))
}
