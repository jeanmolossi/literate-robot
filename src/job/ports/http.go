package ports

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jeanmolossi/literate-robot/src/common/auth"
	"github.com/jeanmolossi/literate-robot/src/common/server/httperr"
	"github.com/jeanmolossi/literate-robot/src/job/app"
	"github.com/jeanmolossi/literate-robot/src/job/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetJobs(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	var appJobs []query.Job

	// simula role diferente
	if user.Role == "recruiter" {
		appJobs, err = h.app.Queries.AllJobs.Handle(r.Context())
	} else {
		appJobs, err = h.app.Queries.AllJobs.Handle(r.Context())
	}

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	jobs := appJobsToResponse(appJobs)
	jobsResp := Jobs{jobs}

	render.Respond(w, r, jobsResp)
}

func (h HttpServer) GetJob(w http.ResponseWriter, r *http.Request, jobID int) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	if user.Role != "recruiter" && user.Role != "admin" {
		httperr.RespondWithSlugError(errors.New("unauthorized"), w, r)
		return
	}

	queryJob, err := h.app.Queries.GetJob.Handle(r.Context(), jobID)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	job := appJobToResponse(queryJob)

	render.Respond(w, r, job)
}

func (h HttpServer) ActivateJob(w http.ResponseWriter, r *http.Request, jobID int) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	if user.Role != "admin" && user.Role != "recruiter" {
		httperr.RespondWithSlugError(errors.New("unauthorized"), w, r)
		return
	}

	err = h.app.Commands.ActivateJob.Handle(r.Context(), jobID)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}

func (h HttpServer) DeactivateJob(w http.ResponseWriter, r *http.Request, jobID int) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	if user.Role != "admin" {
		httperr.Unauthorized("unauthorized", err, w, r)
		return
	}

	err = h.app.Commands.DeactivateJob.Handle(r.Context(), jobID)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}

func appJobToResponse(job query.Job) Job {
	return Job{
		Id:          fmt.Sprintf("%d", job.ID),
		Title:       job.Title,
		Description: job.Description,
		Status:      job.Status,
		Role:        "recruiter",
		City:        &job.City,
		State:       &job.State,
	}
}

func appJobsToResponse(appJobs []query.Job) []Job {
	jobs := make([]Job, len(appJobs))

	for i, appJob := range appJobs {
		jobs[i] = appJobToResponse(appJob)
	}

	return jobs
}
