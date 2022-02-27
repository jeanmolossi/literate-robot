package job_location

import "context"

type Repository interface {
	GetJobCity(ctx context.Context, cityID int) (*JobCity, error)
	GetJobState(ctx context.Context, stateID int) (*JobState, error)
	GetJobLocations(ctx context.Context, jobsIDs []int) ([]JobLocation, error)
}
