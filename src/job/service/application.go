package service

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

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
	db, err := sql.Open("mysql", "root:root@tcp(jobs-db:3306)/jobs")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(10)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(5)

	jobRepository := adapters.NewJobRepository(db)

	return app.Application{
		Commands: app.Commands{
			ActivateJob:   command.NewActivateJobHandler(jobRepository, jobGrpc),
			DeactivateJob: command.NewDeactivateJobHandler(jobRepository, jobGrpc),
		},
		Queries: app.Queries{
			AllJobs: query.NewAllJobsHandler(jobRepository),
			GetJob:  query.NewGetJobHandler(jobRepository),
		},
	}

}
