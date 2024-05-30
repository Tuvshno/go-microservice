package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

/**
The loggingService uses the Decorator design pattern to add additional logging functionality
on top of the PriceFetcher interface. The loggingService struct takes in a PriceFetcher and
delegates the responsibility to fetch the price itself and uses the result of the fetchPrice
to log details about its process.
*/

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
