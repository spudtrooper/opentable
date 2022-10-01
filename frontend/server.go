package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/handlers"
)

func ListenAndServe(ctx context.Context, port int, host string, staticDir string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	client, err := api.MakeExtendedFromFlags(ctx)
	if err != nil {
		return err
	}
	handlers := handlers.CreateHandlers(client)
	handler := handler.CreateHandler(ctx, handlers, handler.CreateHandlerIndexTitle("opentable.com API"))
	if staticDir != "" {
		handler.Handle("/", gorillahandlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(staticDir))))
	}

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		return err
	}

	return nil
}
