package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeanmolossi/literate-robot/src/common/logs"
	"github.com/jeanmolossi/literate-robot/src/common/server"
	"github.com/jeanmolossi/literate-robot/src/job/ports"
)

func main() {
	logs.Init()

	ctx := context.Background()

	app, cleanup := service.NewApplication(ctx)
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
