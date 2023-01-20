package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chinathaip/catfact/model"
)

type Service interface {
	GetCatFact(context.Context) (*model.CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) *CatFactService {
	return &CatFactService{url: url}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*model.CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &model.CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
