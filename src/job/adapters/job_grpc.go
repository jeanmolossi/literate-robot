package adapters

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job"
)

type JobGrpc struct {
	client job.JobServiceClient
}

func NewJobGrpc(client job.JobServiceClient) JobGrpc {
	return JobGrpc{client}
}

func (j JobGrpc) GetJob(ctx context.Context, id string) error {
	return nil
}

func (j JobGrpc) ActivateJob(ctx context.Context, jobID int) error {
	return nil
}

func (j JobGrpc) DeactivateJob(ctx context.Context, jobID int) error {
	return nil
}
