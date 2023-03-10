//lint:file-ignore ST1006

package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)

}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (self *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error){
	response, err := http.Get(self.url);

	if err != nil {
		return nil, err;
	}

	defer response.Body.Close();

	fact := CatFact{};

	if err := json.NewDecoder(response.Body).Decode(&fact); err != nil {
		return nil, err;
	}

	return &fact, nil;
}