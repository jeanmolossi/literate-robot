package app

import (
	"github.com/jeanmolossi/literate-robot/src/job/app/command"
	"github.com/jeanmolossi/literate-robot/src/job/app/query"
)

type (
	Application struct {
		Commands Commands
		Queries  Queries
	}

	Commands struct {
		ActivateJob   command.ActivateJobHandler
		DeactivateJob command.DeactivateJobHandler
	}

	Queries struct {
		AllJobs query.AllJobsHandler
		GetJob  query.GetJobHandler
	}
)
