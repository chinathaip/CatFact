package client

import (
	"context"

	"github.com/chinathaip/catfact/model"
	"github.com/chinathaip/catfact/service"
)

func GetCatFact(ctx context.Context) (*model.CatFact, error) {
	catFactService := service.NewCatFactService("https://catfact.ninja/fact")
	loggingService := service.NewLoggingService(catFactService)
	return loggingService.GetCatFact(ctx)
}
