package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
)

type AllJobsHandler struct {
	readModel job.Repository
}

func NewAllJobsHandler(readModel job.Repository) AllJobsHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return AllJobsHandler{readModel}
}

func (h AllJobsHandler) Handle(ctx context.Context) (jobs []Job, err error) {
	return h.UnmarshalJobFromDatabase(h.readModel.GetAllJobs(ctx))
}

func (h AllJobsHandler) UnmarshalJobFromDatabase(j []job.Job, err error) ([]Job, error) {
	jobs := make([]Job, len(j))

	for i, j := range j {
		jobs[i] = Job{
			ID:          j.JobID(),
			Title:       j.JobTitle(),
			Description: j.JobDescription(),
			Status:      j.JobStatus(),
		}
	}

	return jobs, err
}
