package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opentable/api"
	"github.com/spudtrooper/opentable/handlers"
)

func ListenAndServe(ctx context.Context, client *api.Extended, port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	handlers := handlers.CreateHandlers(client)
	mux, err := handler.CreateHandler(ctx, handlers,
		handler.CreateHandlerPrefix("api"),
		handler.CreateHandlerIndexTitle("unofficial opentable API"),
		handler.CreateHandlerFooterHTML(`Details: <a target="_" href="https://github.com/spudtrooper/opentable">github.com/spudtrooper/opentable</a>`),
		handler.CreateHandlerSourceLinks(true),
		handler.CreateHandlerHandlersFiles([]string{"handlers/handlers.go"}),
		handler.CreateHandlerSourceLinkURIRoot("https://github.com/spudtrooper/opentable/blob/main"),
		handler.CreateHandlerFormatHTML(true),
	)
	if err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
