package cli

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/minimalcli/handler"
	flag "github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/handlers"
)

var (
	failureJSON = flags.String("failure_json", "file to test parsing")
)

func Main(ctx context.Context) error {
	flag.Bool("verbose", false, "global verbose")
	flag.String("term", "", "global term")
	flag.String("uri", "", "global uri")
	flag.String("rest_id", "", `restaurant id, e.g. "kumi-japanese-restaurant-and-bar-nyc-new-york"`)
	flag.Int("start_page", 0, "global start page")
	flag.Int("threads", 0, "global threads")
	flag.Duration("sleep", 0, "global sleep")
	flag.Bool("debug_failures", false, "global debug_failres")

	adp := handler.NewCLIAdapter()
	adp.BindToGlobalFlags()

	client, err := api.MakeExtendedFromFlags(ctx)
	if err != nil {
		return err
	}

	handlers := handlers.CreateHandlers(client)
	app := handler.CreateCLI(adp, handlers...)

	app.Register("TestFailedJSON", func(context.Context) error {
		requireStringFlag(failureJSON, "failure_json")
		var payload interface{}
		log.Printf("testing failed json %q", *failureJSON)
		b, err := ioutil.ReadFile(*failureJSON)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &payload); err != nil {
			return err
		}
		return nil
	})

	app.AddPostRunHook(func(ctx context.Context) error {
		log.Printf("\n\nStats\n%s", client.StatsString())
		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}

func requireStringFlag(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}
