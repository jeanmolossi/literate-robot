package adapters

import (
	"context"
	"database/sql"

	"github.com/jeanmolossi/literate-robot/src/job/domain/job"
)

type dbJob struct {
	ID          int
	Title       string
	Description string
	Status      string
	CityID      int
	StateID     int
	UserID      int
}

type JobRepository struct {
	client *sql.DB
}

func NewJobRepository(client *sql.DB) JobRepository {
	return JobRepository{client}
}

func (j JobRepository) GetAllJobs(ctx context.Context) ([]job.Job, error) {
	rows, err := j.client.QueryContext(ctx, "SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}

	var jobs []job.Job
	for rows.Next() {
		var bdJob dbJob
		err := rows.Scan(&bdJob.ID, &bdJob.Title, &bdJob.Description, &bdJob.Status, &bdJob.CityID, &bdJob.StateID, &bdJob.UserID)
		if err != nil {
			return nil, err
		}

		domainJob, err := job.NewJob(bdJob.ID, bdJob.Title, bdJob.Description, bdJob.Status, bdJob.UserID, bdJob.CityID, bdJob.StateID)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, *domainJob)
	}

	return jobs, nil
}

func (j JobRepository) GetJob(ctx context.Context, jobId int) (*job.Job, error) {
	row := j.client.QueryRowContext(ctx, "SELECT * FROM jobs WHERE job_id = ?", jobId)

	var bdJob dbJob
	if err := row.Scan(&bdJob.ID, &bdJob.Title, &bdJob.Description, &bdJob.Status, &bdJob.CityID, &bdJob.StateID, &bdJob.UserID); err != nil {
		return nil, err
	}

	domainJob, err := job.NewJob(bdJob.ID, bdJob.Title, bdJob.Description, bdJob.Status, bdJob.UserID, bdJob.CityID, bdJob.StateID)
	if err != nil {
		return nil, err
	}

	return domainJob, nil
}

func (j JobRepository) UpdateJob(ctx context.Context, jobID int, updateFunc func(*job.Job) (*job.Job, error)) error {
	_job, err := j.GetJob(ctx, jobID)
	if err != nil {
		return err
	}

	_updatedJob, err := updateFunc(_job)
	if err != nil {
		return err
	}

	row := j.client.QueryRowContext(ctx,
		"UPDATE jobs SET title = ?, description = ?, status = ?, city_id = ?, state_id = ?, user_id = ? WHERE job_id = ?",
		_updatedJob.JobTitle(), _updatedJob.JobDescription(), _updatedJob.JobStatus(), _updatedJob.JobCityID(), _updatedJob.JobStateID(), _updatedJob.JobUserID(), _updatedJob.JobID(),
	)

	return row.Err()
}
