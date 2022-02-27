package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
)

type GetJobHandler struct {
	readModel job.Repository
	client    JobLocationService
}

func NewGetJobHandler(readModel job.Repository, client JobLocationService) GetJobHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return GetJobHandler{readModel, client}
}

func (g GetJobHandler) Handle(ctx context.Context, jobId int) (_job Job, err error) {
	j, err := g.readModel.GetJob(ctx, jobId)
	if err != nil {
		panic(err)
	}

	jc, err := g.client.GetJobCity(ctx, j.JobCityID())
	if err != nil {
		panic(err)
	}
	js, err := g.client.GetJobState(ctx, j.JobStateID())
	if err != nil {
		panic(err)
	}

	jLoc, err := job.NewJobLocation(jc.City, js.State)
	if err != nil {
		panic(err)
	}

	_job, err = g.UnmarshalJobFromDatabase(j, err)
	_job.City = jLoc.City()
	_job.State = jLoc.State()

	return _job, err
}

func (g GetJobHandler) UnmarshalJobFromDatabase(j *job.Job, err error) (Job, error) {

	return Job{
		ID:          j.JobID(),
		Title:       j.JobTitle(),
		Description: j.JobDescription(),
		Status:      j.JobStatus(),
	}, err
}
