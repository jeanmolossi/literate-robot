package query

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job_location/domain/job_location"
)

type GetJobCityHandler struct {
	readModel job_location.Repository
}

func NewGetJobCityHandler(readModel job_location.Repository) GetJobCityHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return GetJobCityHandler{readModel}
}

func (g GetJobCityHandler) Handle(ctx context.Context, cityID int) (JobCity, error) {
	return g.UnmarshalFromDatabase(g.readModel.GetJobCity(ctx, cityID))
}

func (g GetJobCityHandler) UnmarshalFromDatabase(j *job_location.JobCity, err error) (JobCity, error) {
	return JobCity{
		CityID: j.CityID(),
		City:   j.City(),
	}, err
}
