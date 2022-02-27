package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job_location/domain/job_location"
)

type GetJobsLocationsHandler struct {
	readModel job_location.Repository
}

func NewGetJobsLocationsHandler(readModel job_location.Repository) GetJobsLocationsHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return GetJobsLocationsHandler{readModel}
}

func (g GetJobsLocationsHandler) Handle(ctx context.Context, jobsIDs []int) ([]JobLocation, error) {
	return g.UnmarshalFromDatabase(g.readModel.GetJobLocations(ctx, jobsIDs))
}

func (g GetJobsLocationsHandler) UnmarshalFromDatabase(jobs []job_location.JobLocation, err error) ([]JobLocation, error) {
	var _jobs []JobLocation
	for _, j := range jobs {
		_jobs = append(_jobs, JobLocation{
			JobID: j.JobID,
			JobCity: JobCity{
				CityID: j.JobCity.CityID(),
				City:   j.JobCity.City(),
			},
			JobState: JobState{
				StateID: j.JobState.StateID(),
				State:   j.JobState.State(),
			},
		})
	}

	return _jobs, nil
}
