package apihandler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
)

func CreateHandler(ctx context.Context, staticDir string, hostPort string, hs ...Handler) http.Handler {
	mux := http.NewServeMux()
	if staticDir != "" {
		mux.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(staticDir))))
	}

	log.Printf("Initializing server...")

	handleFunc := func(route string, fn func(w http.ResponseWriter, req *http.Request)) {
		log.Printf("  route %s %s%s", route, hostPort, route)
		mux.HandleFunc(route, fn)
	}

	for _, h := range hs {
		h := h.(*handler)
		if h.cliOnly {
			continue
		}
		route := fmt.Sprintf("/api/%s", strings.ToLower(h.name))
		handleFunc(route, func(w http.ResponseWriter, req *http.Request) {
			evalCtx := &serverEvalContext{
				ctx: ctx,
				w:   w,
				req: req,
			}
			res, err := h.fn(evalCtx)
			if err != nil {
				respondWithError(w, req, err)
				return
			}
			respondWithJSON(req, w, res)
		})
	}

	return mux
}
