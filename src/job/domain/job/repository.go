package job

import "context"

type Repository interface {
	GetJob(ctx context.Context, jobId int) (*Job, error)
}
