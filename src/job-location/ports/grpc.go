package ports

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job_location"
	"github.com/jeanmolossi/literate-robot/src/job_location/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) GetJobCity(ctx context.Context, in *job_location.GetJobCityRequest) (*job_location.GetJobCityResponse, error) {
	jobCity, err := g.app.Queries.GetJobCity.Handle(ctx, int(in.CityId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &job_location.GetJobCityResponse{
		City:   jobCity.City,
		CityId: int32(jobCity.CityID),
	}, err
}

func (g GrpcServer) GetJobState(ctx context.Context, in *job_location.GetJobStateRequest) (*job_location.GetJobStateResponse, error) {
	jobState, err := g.app.Queries.GetJobState.Handle(ctx, int(in.StateId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &job_location.GetJobStateResponse{
		State:   jobState.State,
		StateId: int32(jobState.StateID),
	}, err
}

func (g GrpcServer) GetJobLocations(ctx context.Context, in *job_location.GetJobLocationsRequest) (*job_location.GetJobLocationsResponse, error) {
	ids := make([]int, len(in.GetJobIds()))
	for i, id := range in.JobIds {
		ids[i] = int(id)
	}

	jobsLocations, err := g.app.Queries.GetJobsLocations.Handle(ctx, ids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var jobLocations []*job_location.SingleJobLocationResponse
	for _, jobLocation := range jobsLocations {
		jobLocations = append(jobLocations, &job_location.SingleJobLocationResponse{
			City:    jobLocation.City,
			CityId:  int32(jobLocation.CityID),
			State:   jobLocation.State,
			StateId: int32(jobLocation.StateID),
			JobId:   int32(jobLocation.JobID),
		})
	}

	return &job_location.GetJobLocationsResponse{
		JobLocations: jobLocations,
	}, err
}
