package adapters

import (
	"context"
	"database/sql"

	"github.com/jeanmolossi/literate-robot/src/job_location/domain/job_location"
)

type JobCity struct {
	CityID int
	City   string
}

type JobState struct {
	StateID int
	State   string
}

type Repository struct {
	client *sql.DB
}

func NewJobLocationRepository(client *sql.DB) *Repository {
	return &Repository{client}
}

func (j Repository) GetJobCity(ctx context.Context, cityID int) (*job_location.JobCity, error) {
	row := j.client.QueryRowContext(ctx, "SELECT city_id, city FROM job_cities WHERE city_id = ?", cityID)

	var jobCity JobCity
	err := row.Scan(&jobCity.CityID, &jobCity.City)
	if err != nil {
		panic(err)
	}

	return job_location.NewJobCity(jobCity.CityID, jobCity.City), nil
}

func (j Repository) GetJobState(ctx context.Context, stateID int) (*job_location.JobState, error) {
	row := j.client.QueryRowContext(ctx, "SELECT state_id, state FROM job_states WHERE state_id = ?", stateID)

	var jobState JobState
	err := row.Scan(&jobState.StateID, &jobState.State)
	if err != nil {
		panic(err)
	}

	return job_location.NewJobState(jobState.StateID, jobState.State), nil
}

func (j Repository) GetJobLocations(ctx context.Context, jobsIDs []int) ([]job_location.JobLocation, error) {
	in, jobsIds := InConditionWithInt(jobsIDs)

	rows, err := j.client.QueryContext(ctx,
		`SELECT j.job_id, jc.city_id, jc.city, js.state_id, js.state FROM jobs AS j`+
			` INNER JOIN job_cities jc ON j.city_id = jc.city_id`+
			` INNER JOIN job_states js ON j.state_id = js.state_id WHERE j.job_id IN `+in,
		jobsIds...,
	)
	if err != nil {
		panic(err)
	}

	var jobLocations []job_location.JobLocation
	for rows.Next() {
		var jobID int
		var jobCity JobCity
		var jobState JobState
		err := rows.Scan(&jobID, &jobCity.CityID, &jobCity.City, &jobState.StateID, &jobState.State)
		if err != nil {
			panic(err)
		}

		jobLocations = append(jobLocations, *job_location.NewJobLocation(
			jobID,
			job_location.NewJobCity(jobCity.CityID, jobCity.City),
			job_location.NewJobState(jobState.StateID, jobState.State)),
		)
	}

	return jobLocations, nil
}

func InConditionWithInt(params []int) (string, []interface{}) {
	in := "("
	var values []interface{}
	for _, param := range params {
		in += "?,"
		values = append(values, param)
	}
	in = in[:len(in)-1] + ")"
	return in, values
}
