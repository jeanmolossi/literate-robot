package job

import "errors"

type Job struct {
	id          string
	role        string
	title       string
	description string
	status      string
}

func NewJob(id string, role string, title string, description string, status string) (*Job, error) {
	if id == "" {
		return nil, errors.New("empty job id")
	}

	if role == "" {
		return nil, errors.New("empty job role")
	}

	if title == "" {
		return nil, errors.New("empty job title")
	}

	return &Job{
		id:          id,
		role:        role,
		title:       title,
		description: description,
		status:      status,
	}, nil
}
