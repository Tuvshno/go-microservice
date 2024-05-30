package main

import (
	"context"
	"fmt"
	"time"
)

/*
	service has the contains the implementation of the PriceFetcher Interface which implements FetchPrice.
	service also has dummy mock data that the FetchPrice calls to as a mock API call.
*/

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

// priceMocks holds dummy data of prices.
var priceMocks = map[string]float64{
	"BTC":  20_000.0,
	"ETH":  200.0,
	"RAND": 100_000.0,
}

// MockPriceFetcher is a dummy function that mimics a real API call to the price API.
func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic the HTTP round trip
	time.Sleep(100 * time.Millisecond)

	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not suppported", ticker)
	}
	return price, nil
}
