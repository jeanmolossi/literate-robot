package app

import "github.com/jeanmolossi/literate-robot/src/job/app/query"

type (
	Application struct {
		Commands Commands
		Queries  Queries
	}

	Commands struct{}

	Queries struct {
		AllJobs query.AllJobsHandler
	}
)
