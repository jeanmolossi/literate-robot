package ports

import (
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

func appJobsToResponse(appJobs []query.Job) []Job {
	jobs := make([]Job, len(appJobs))

	for i, appJob := range appJobs {
		jobs[i] = Job{
			Id:          fmt.Sprintf("%d", appJob.ID),
			Title:       appJob.Title,
			Description: appJob.Description,
			Status:      appJob.Status,
			Role:        "recruiter",
		}
	}

	return jobs
}
