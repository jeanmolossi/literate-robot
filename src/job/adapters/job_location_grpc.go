package adapters

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job_location"
)

type JobLocationGrpc struct {
	client job_location.JobLocationServiceClient
}

func NewJobLocationGrpc(client job_location.JobLocationServiceClient) JobLocationGrpc {
	return JobLocationGrpc{client}
}

func (j JobLocationGrpc) GetJobCity(ctx context.Context, cityID int) (*job_location.GetJobCityResponse, error) {
	jobCity, err := j.client.GetJobCity(ctx, &job_location.GetJobCityRequest{
		CityId: int32(cityID),
	})

	if err != nil {
		return nil, err
	}

	return jobCity, nil
}

func (j JobLocationGrpc) GetJobState(ctx context.Context, stateID int) (*job_location.GetJobStateResponse, error) {
	jobState, err := j.client.GetJobState(ctx, &job_location.GetJobStateRequest{
		StateId: int32(stateID),
	})

	if err != nil {
		return nil, err
	}

	return jobState, nil
}

func (j JobLocationGrpc) GetJobLocations(ctx context.Context, jobsIDs []int) (*job_location.GetJobLocationsResponse, error) {
	ids := make([]int32, len(jobsIDs))

	for i, id := range jobsIDs {
		ids[i] = int32(id)
	}

	jobsLocations, err := j.client.GetJobLocations(ctx, &job_location.GetJobLocationsRequest{
		JobIds: ids,
	})

	if err != nil {
		return nil, err
	}

	return jobsLocations, nil
}
