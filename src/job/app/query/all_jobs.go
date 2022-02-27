package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
)

type AllJobsHandler struct {
	readModel job.Repository
	client    JobLocationService
}

func NewAllJobsHandler(readModel job.Repository, client JobLocationService) AllJobsHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return AllJobsHandler{readModel, client}
}

func (h AllJobsHandler) Handle(ctx context.Context) (jobs []Job, err error) {
	return h.UnmarshalJobFromDatabase(h.readModel.GetAllJobs(ctx))
}

func (h AllJobsHandler) UnmarshalJobFromDatabase(j []job.Job, _ error) ([]Job, error) {
	jobs := make([]Job, len(j))

	ctx := context.Background()

	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].JobID()
	}

	jobLocations, err := h.client.GetJobLocations(ctx, ids)

	for i, j := range j {
		var jobLocation job.JobLocation
		for _, jl := range jobLocations.JobLocations {
			if int(jl.JobId) == j.JobID() {
				nJobLocation, err := job.NewJobLocation(jl.City, jl.State)
				jobLocation = *nJobLocation
				if err != nil {
					panic(err)
				}
				break
			}
		}

		jobs[i] = Job{
			ID:          j.JobID(),
			Title:       j.JobTitle(),
			Description: j.JobDescription(),
			Status:      j.JobStatus(),
			City:        jobLocation.City(),
			State:       jobLocation.State(),
		}
	}

	return jobs, err
}
