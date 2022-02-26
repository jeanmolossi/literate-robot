package command

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeactivateJob struct {
	JobID int
}

type DeactivateJobHandler struct {
	repository job.Repository
	jobService JobService
}

func NewDeactivateJobHandler(repo job.Repository, jobService JobService) DeactivateJobHandler {
	if repo == nil {
		panic("repository is nil")
	}

	if jobService == nil {
		panic("jobService is nil")
	}

	return DeactivateJobHandler{repository: repo, jobService: jobService}
}

func (a DeactivateJobHandler) Handle(ctx context.Context, jobID int) (err error) {
	_job, err := a.repository.GetJob(ctx, jobID)
	if err != nil {
		return err
	}

	if !_job.IsActive() {
		return status.Error(codes.FailedPrecondition, "job is not in draft status")
	}

	err = a.repository.UpdateJob(ctx, jobID, func(newJob *job.Job) (*job.Job, error) {
		newJob.Deactivate()
		return newJob, nil
	})

	if err != nil {
		return err
	}

	return a.jobService.DeactivateJob(ctx, jobID)
}
