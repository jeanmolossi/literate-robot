package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jeanmolossi/literate-robot/src/common/genproto/job"
	"github.com/jeanmolossi/literate-robot/src/common/logs"
	"github.com/jeanmolossi/literate-robot/src/common/server"
	"github.com/jeanmolossi/literate-robot/src/job/ports"
	"github.com/jeanmolossi/literate-robot/src/job/service"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()

	ctx := context.Background()

	app, cleanup := service.NewApplication(ctx)
	defer cleanup()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		server.RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(app), router)
		})
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			svc := ports.NewGrpcServer(app)
			job.RegisterJobServiceServer(server, svc)
		})
	default:
		panic("server type not supported")

	}
}
