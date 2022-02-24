package service

import (
	"context"

	"cloud.google.com/go/firestore"
	grpcClient "github.com/jeanmolossi/literate-robot/src/common/client"
	"github.com/jeanmolossi/literate-robot/src/job/adapters"
	"github.com/jeanmolossi/literate-robot/src/job/app"
	"github.com/jeanmolossi/literate-robot/src/job/app/command"
	"github.com/jeanmolossi/literate-robot/src/job/app/query"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	jobClient, closeJobClient, err := grpcClient.NewJobClient()
	if err != nil {
		panic(err)
	}

	jobGrpc := adapters.NewJobGrpc(jobClient)

	return newApplication(ctx, jobGrpc), func() {
		_ = closeJobClient()
	}
}

func newApplication(ctx context.Context, jobGrpc command.JobService) app.Application {
	client, err := firestore.NewClient(ctx, "literate-robot")
	if err != nil {
		panic(err)
	}

	jobRepository := adapters.NewJobRepository(client)

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			AllJobs: query.NewAllJobsHandler(jobRepository),
		},
	}

}
