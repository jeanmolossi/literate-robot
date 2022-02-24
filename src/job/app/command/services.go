package command

import "context"

type JobService interface {
	GetJob(ctx context.Context, id string) error
}
