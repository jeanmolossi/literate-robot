package adapters

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/jeanmolossi/literate-robot/src/job/app/query"
)

type JobRepository struct {
	client *firestore.Client
}

func NewJobRepository(client *firestore.Client) JobRepository {
	return JobRepository{client}
}

func (j JobRepository) AllJobs(ctx context.Context) ([]query.Job, error) {
	createdAt := time.Now()

	return []query.Job{
		{
			ID:          "1",
			Description: "Job 1 description",
			Status:      "pending",
			Role:        "developer",
			CreatedAt:   &createdAt,
			Title:       "Job 1 title",
			Location:    nullishString("Job 1 location"),
			Company:     nullishInt(149),
			Salary:      nullishString("Job 1 salary"),
		},
		{
			ID:          "2",
			Description: "Job 2 description",
			Status:      "inspection",
			Role:        "developer",
			CreatedAt:   &createdAt,
			Title:       "Job 2 title",
			Location:    nullishString("Job 2 location"),
			Company:     nullishInt(149),
			Salary:      nullishString("Job 2 salary"),
		},
	}, nil
}

func nullishInt(i int) *int {
	if i == 0 {
		return nil
	}

	return &i
}

func nullishString(txt string) *string {
	if txt == "" {
		return nil
	}

	return &txt
}
