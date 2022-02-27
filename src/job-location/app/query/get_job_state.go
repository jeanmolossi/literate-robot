package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job_location/domain/job_location"
)

type GetJobStateHandler struct {
	readModel job_location.Repository
}

func NewGetJobStateHandler(readModel job_location.Repository) GetJobStateHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return GetJobStateHandler{readModel}
}

func (g GetJobStateHandler) Handle(ctx context.Context, stateID int) (JobState, error) {
	return g.UnmarshalFromDatabase(g.readModel.GetJobState(ctx, stateID))
}

func (g GetJobStateHandler) UnmarshalFromDatabase(j *job_location.JobState, err error) (JobState, error) {
	return JobState{
		StateID: j.StateID(),
		State:   j.State(),
	}, err
}
