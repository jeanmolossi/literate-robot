package job

import "context"

type Repository interface {
	GetJob(ctx context.Context, jobId int) (*Job, error)
	GetAllJobs(ctx context.Context) ([]Job, error)

	UpdateJob(ctx context.Context, jobID int, updateFunc func(j *Job) (*Job, error)) error
}
