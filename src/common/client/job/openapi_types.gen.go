// Package job provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package job

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Job defines model for Job.
type Job struct {
	Company     *int       `json:"company,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	Description string     `json:"description"`
	Id          string     `json:"id"`
	Location    *string    `json:"location,omitempty"`
	Role        string     `json:"role"`
	Salary      *string    `json:"salary,omitempty"`
	Status      string     `json:"status"`
	Title       string     `json:"title"`
}

// Jobs defines model for Jobs.
type Jobs struct {
	Jobs []Job `json:"jobs"`
}
