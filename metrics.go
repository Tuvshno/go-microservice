package main

import (
	"context"
	"fmt"
)

/*
metrics has the implementation of the metrics middleware and acts as a decorator on the PriceFetcher interface.
*/

// metricService is a decorator on top of the PriceFetcher interface
type metricService struct {
	next PriceFetcher
}

// NewMetricService returns a PriceFetcher interface with the decorator Metric Service on top
func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

// FetchPrice implements the FetchPrice function of the PriceFetcher interface pushes metrics to the storage
func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	// Metrics storage. Push to prometheus (gauge, counters)
	fmt.Println("Pushing to Metrics Storage System")
	return s.next.FetchPrice(ctx, ticker)
}
