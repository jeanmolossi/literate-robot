package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job_location"
)

type Job struct {
	ID          int
	Title       string
	Description string
	Status      string
	City        string
	State       string
}

type JobCity struct {
	CityID int
	City   string
}

type JobState struct {
	State string
}

type JobLocationService interface {
	GetJobCity(ctx context.Context, cityID int) (*job_location.GetJobCityResponse, error)
	GetJobState(ctx context.Context, stateID int) (*job_location.GetJobStateResponse, error)
	GetJobLocations(ctx context.Context, jobsIDs []int) (*job_location.GetJobLocationsResponse, error)
}
