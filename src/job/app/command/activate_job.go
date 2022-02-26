package command

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActivateJob struct {
	JobID int
}

type ActivateJobHandler struct {
	repository job.Repository
	jobService JobService
}

func NewActivateJobHandler(repo job.Repository, jobService JobService) ActivateJobHandler {
	if repo == nil {
		panic("repository is nil")
	}

	if jobService == nil {
		panic("jobService is nil")
	}

	return ActivateJobHandler{repository: repo, jobService: jobService}
}

func (a ActivateJobHandler) Handle(ctx context.Context, jobID int) (err error) {
	_job, err := a.repository.GetJob(ctx, jobID)
	if err != nil {
		return err
	}

	if _job.IsActive() {
		return status.Error(codes.FailedPrecondition, "job is not in draft status")
	}

	err = a.repository.UpdateJob(ctx, jobID, func(newJob *job.Job) (*job.Job, error) {
		newJob.Activate()
		return newJob, nil
	})
	if err != nil {
		return err
	}

	return a.jobService.ActivateJob(ctx, jobID)
}
