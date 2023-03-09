package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (self *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {

	defer func(start time.Time){
		fmt.Printf("fact=%s, err=%v , it took %v \n", fact.Fact, err, time.Since(start))
	}(time.Now());


	return self.next.GetCatFact(ctx);
}