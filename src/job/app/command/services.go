package command

import "context"

type JobService interface {
	ActivateJob(ctx context.Context, jobID int) error
	DeactivateJob(ctx context.Context, jobID int) error
}
