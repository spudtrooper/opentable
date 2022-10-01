// Package apihandler generates a CLI and frontend API server to
// access a third party API.
package apihandler

import (
	"context"
	"fmt"
	"time"

	goutiljson "github.com/spudtrooper/goutil/json"
	minimalcli "github.com/spudtrooper/minimalcli/app"
)

type CLIAdapter struct {
	stringFlags map[string]*string
	boolFlags   map[string]*bool
	intFlags    map[string]*int
	durFlags    map[string]*time.Duration
}

func NewCLIAdapter() *CLIAdapter {
	return &CLIAdapter{
		stringFlags: make(map[string]*string),
		boolFlags:   make(map[string]*bool),
		intFlags:    make(map[string]*int),
		durFlags:    make(map[string]*time.Duration),
	}
}

func (c *CLIAdapter) BindStringFlag(name string, flag *string)          { c.stringFlags[name] = flag }
func (c *CLIAdapter) BindBoolFlag(name string, flag *bool)              { c.boolFlags[name] = flag }
func (c *CLIAdapter) BindIntFlag(name string, flag *int)                { c.intFlags[name] = flag }
func (c *CLIAdapter) BindDurationFlag(name string, flag *time.Duration) { c.durFlags[name] = flag }

func CreateApp(clia *CLIAdapter, hs ...Handler) *minimalcli.App {
	app := minimalcli.Make()
	app.Init()

	for _, h := range hs {
		h := h.(*handler)
		app.Register(h.name, func(ctx context.Context) error {
			evalCtx := &cliEvalContext{
				ctx:         ctx,
				stringFlags: clia.stringFlags,
				boolFlags:   clia.boolFlags,
				intFlags:    clia.intFlags,
				durFlags:    clia.durFlags,
			}
			res, err := h.fn(evalCtx)
			if err != nil {
				return err
			}
			fmt.Printf("%s: %s", h.name, mustFormatString(res))
			return nil
		})
	}

	return app
}

func mustFormatString(x interface{}) string {
	return goutiljson.MustColorMarshal(x)
}
