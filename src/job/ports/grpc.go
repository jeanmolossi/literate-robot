package ports

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
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
		if job.ID == jobId {
			grpcJob.Id = job.ID
			grpcJob.Name = job.Title
			grpcJob.Description = job.Description
			grpcJob.CreatedAt = &timestamp.Timestamp{Seconds: job.CreatedAt.Unix()}
			grpcJob.UpdatedAt = &timestamp.Timestamp{Seconds: time.Now().Unix()}
		}
	}

	return grpcJob, nil
}
