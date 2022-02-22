package service

import (
	"context"

	"github.com/jeanmolossi/literate-robot/src/job/app"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	jobClient, closeJobClient, err := grpcClient.NewJobClient()
	if err != nil {
		panic(err)
	}
	return app.Application{}, func() {}
}
