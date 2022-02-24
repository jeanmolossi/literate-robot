package adapters

import (
	"context"
	"database/sql"

	"github.com/jeanmolossi/literate-robot/src/job/app/query"

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

func (j JobRepository) AllJobs(ctx context.Context) ([]query.Job, error) {
	rows, err := j.client.QueryContext(ctx, "SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}

	var jobs []query.Job
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

		jobs = append(jobs, j.domainJobToQueryJob(domainJob))
	}

	return jobs, nil
}

func (j JobRepository) domainJobToQueryJob(dJob *job.Job) query.Job {
	return query.Job{
		ID:          dJob.JobID(),
		Title:       dJob.JobTitle(),
		Description: dJob.JobDescription(),
		Status:      dJob.JobStatus(),
	}
}
