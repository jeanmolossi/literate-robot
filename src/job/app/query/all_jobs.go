package query

import "context"

type AllJobsHandler struct {
	readModel AllJobsReadModel
}

func NewAllJobsHandler(readModel AllJobsReadModel) AllJobsHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return AllJobsHandler{readModel}
}

type AllJobsReadModel interface {
	AllJobs(ctx context.Context) ([]Job, error)
}

func (h AllJobsHandler) Handle(ctx context.Context) (jobs []Job, err error) {
	return h.readModel.AllJobs(ctx)
}
