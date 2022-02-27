package app

import "github.com/jeanmolossi/literate-robot/src/job_location/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}

type Queries struct {
	GetJobCity       query.GetJobCityHandler
	GetJobState      query.GetJobStateHandler
	GetJobsLocations query.GetJobsLocationsHandler
}
