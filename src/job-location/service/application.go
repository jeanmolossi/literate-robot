package service

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	grpcClient "github.com/jeanmolossi/literate-robot/src/common/client"
	"github.com/jeanmolossi/literate-robot/src/job_location/adapters"
	"github.com/jeanmolossi/literate-robot/src/job_location/app"
	"github.com/jeanmolossi/literate-robot/src/job_location/app/command"
	"github.com/jeanmolossi/literate-robot/src/job_location/app/query"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	jobLocationClient, closeJobLocationClient, err := grpcClient.NewJobLocationClient()
	if err != nil {
		panic(err)
	}

	return newApplication(ctx, jobLocationClient), func() {
		_ = closeJobLocationClient()
	}
}

func newApplication(ctx context.Context, jobLocationGrpc command.JobLocationService) app.Application {
	db, err := sql.Open("mysql", "root:root@tcp(jobs-db:3306)/jobs")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(10)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(5)

	jobLocationRepository := adapters.NewJobLocationRepository(db)

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetJobCity:       query.NewGetJobCityHandler(jobLocationRepository),
			GetJobState:      query.NewGetJobStateHandler(jobLocationRepository),
			GetJobsLocations: query.NewGetJobsLocationsHandler(jobLocationRepository),
		},
	}
}
