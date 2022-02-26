package job

import "errors"

type Job struct {
	job_id      int
	title       string
	description string
	status      string
	user_id     int
	city_id     int
	state_id    int
}

func NewJob(job_id int, title, description, status string, user_id, city_id, state_id int) (*Job, error) {
	if job_id == 0 {
		return nil, errors.New("empty job id")
	}

	if title == "" {
		return nil, errors.New("empty job title")
	}

	if description == "" {
		return nil, errors.New("empty job description")
	}

	if status == "" {
		return nil, errors.New("empty job status")
	}

	if user_id == 0 {
		return nil, errors.New("empty user id")
	}

	if city_id == 0 {
		return nil, errors.New("empty city id")
	}

	if state_id == 0 {
		return nil, errors.New("empty state id")
	}

	return &Job{
		job_id:      job_id,
		title:       title,
		description: description,
		status:      status,
		city_id:     city_id,
		state_id:    state_id,
		user_id:     user_id,
	}, nil
}

func (j Job) JobID() int {
	return j.job_id
}

func (j Job) JobTitle() string {
	return j.title
}

func (j Job) JobDescription() string {
	return j.description
}

func (j Job) JobStatus() string {
	return j.status
}

func (j Job) JobUserID() int {
	return j.user_id
}

func (j Job) JobCityID() int {
	return j.city_id
}

func (j Job) JobStateID() int {
	return j.state_id
}

func (j Job) IsActive() bool {
	return j.status == "A"
}

func (j *Job) Activate() {
	j.status = "A"
}

func (j *Job) Deactivate() {
	j.status = "I"
}
