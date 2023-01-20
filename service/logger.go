package service

import (
	"context"
	"log"
	"time"

	"github.com/chinathaip/catfact/model"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) *LoggingService {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (*model.CatFact, error) {
	defer func(startTime time.Time) {
		log.Printf("GetCatFact took %s", time.Since(startTime))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
