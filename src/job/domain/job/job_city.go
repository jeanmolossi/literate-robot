package job

import "errors"

type JobLocation struct {
	city  string
	state string
}

func NewJobLocation(city, state string) (*JobLocation, error) {
	if city == "" {
		return nil, errors.New("empty city")
	}

	if state == "" {
		return nil, errors.New("empty state")
	}

	return &JobLocation{
		city:  city,
		state: state,
	}, nil
}

func (j *JobLocation) City() string {
	return j.city
}

func (j *JobLocation) State() string {
	return j.state
}
