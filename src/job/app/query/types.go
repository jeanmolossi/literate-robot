package query

import "time"

type Job struct {
	ID          string
	Role        string
	Title       string
	Description string
	Status      string
	Location    *string
	Company     *int
	Salary      *string
	CreatedAt   *time.Time
}
