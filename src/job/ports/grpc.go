package ports

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jeanmolossi/literate-robot/src/common/genproto/job"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jeanmolossi/literate-robot/src/job/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) GetJob(ctx context.Context, jobRequest *job.GetJobRequest) (*job.Job, error) {
	jobs, err := g.app.Queries.AllJobs.Handle(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	jobId := jobRequest.GetId()
	grpcJob := &job.Job{}
	for _, job := range jobs {
		if job.ID == int(jobId) {
			grpcJob.JobId = int32(job.ID)
			grpcJob.Title = job.Title
			grpcJob.Description = job.Description
			grpcJob.Status = job.Status
		}
	}

	return grpcJob, nil
}

func (g GrpcServer) ActivateJob(ctx context.Context, jobID *job.ActivateJobRequest) (*empty.Empty, error) {
	err := g.app.Commands.ActivateJob.Handle(ctx, int(jobID.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) DeactivateJob(ctx context.Context, jobID *job.DeactivateJobRequest) (*empty.Empty, error) {
	err := g.app.Commands.DeactivateJob.Handle(ctx, int(jobID.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
