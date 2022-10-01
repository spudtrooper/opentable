package clinew

import (
	"context"
	"log"

	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/apihandler"
	"github.com/spudtrooper/opentable/handlers"
)

var (
	verbose       = flags.Bool("verbose", "global verbose")
	term          = flags.String("term", "global term")
	uri           = flags.String("uri", "global uri")
	restID        = flags.String("rest_id", `restaurant id, e.g. "kumi-japanese-restaurant-and-bar-nyc-new-york"`)
	startPage     = flags.Int("start_page", "global start page")
	threads       = flags.Int("threads", "global threads")
	sleep         = flags.Duration("sleep", "global sleep")
	debugFailures = flags.Bool("debug_failures", "global debug_failres")
	failureJSON   = flags.String("failure_json", "file to test parsing")
)

func Main(ctx context.Context) error {
	adp := apihandler.NewCLIAdapter()
	adp.BindStringFlag("term", term)
	adp.BindStringFlag("uri", uri)
	adp.BindBoolFlag("verbose", verbose)
	adp.BindBoolFlag("debug_failures", debugFailures)
	adp.BindIntFlag("start_page", startPage)
	adp.BindIntFlag("threads", threads)
	adp.BindStringFlag("rest_id", restID)
	adp.BindDurationFlag("sleep", sleep)

	client, err := api.MakeExtendedFromFlags(ctx)
	if err != nil {
		return err
	}

	handlers := handlers.CreateHandlers(client)
	app := apihandler.CreateApp(adp, handlers...)

	app.AddPostRunHook(func(ctx context.Context) error {
		log.Printf("\n\nStats\n%s", client.StatsString())
		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
