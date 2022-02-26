package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
)

type GetJobHandler struct {
	readModel job.Repository
}

func NewGetJobHandler(readModel job.Repository) GetJobHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return GetJobHandler{readModel}
}

func (g GetJobHandler) Handle(ctx context.Context, jobId int) (job Job, err error) {
	return g.UnmarshalJobFromDatabase(g.readModel.GetJob(ctx, jobId))
}

func (g GetJobHandler) UnmarshalJobFromDatabase(j *job.Job, err error) (Job, error) {
	return Job{
		ID:          j.JobID(),
		Title:       j.JobTitle(),
		Description: j.JobDescription(),
		Status:      j.JobStatus(),
	}, err
}
