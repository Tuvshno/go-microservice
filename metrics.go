package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	// Metrics storage. Push to prometheus (gauge, counters)
	fmt.Println("Pushing to Metrics Storage System")
	return s.next.FetchPrice(ctx, ticker)
}
