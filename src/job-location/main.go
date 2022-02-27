package main

import (
	"context"
	"os"
	"strings"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job_location"
	"github.com/jeanmolossi/literate-robot/src/common/logs"
	"github.com/jeanmolossi/literate-robot/src/common/server"
	"github.com/jeanmolossi/literate-robot/src/job_location/ports"
	"github.com/jeanmolossi/literate-robot/src/job_location/service"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()

	ctx := context.Background()

	app, cleanup := service.NewApplication(ctx)
	defer cleanup()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			svc := ports.NewGrpcServer(app)
			job_location.RegisterJobLocationServiceServer(server, svc)
		})
	default:
		panic("server type not supported")
	}
}
