package main

import (
	"context"
	"encoding/json"
	"main/types"
	"math/rand"
	"net/http"
)

/**
api holds the JSON transport for fetching prices from the PriceFetcher service
*/

// APIFunc is a type alias that takes a context, response writer, a request, and returns an error
type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

// JSONAPIServer is a struct that holds the JSON API server information including port and service
type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

// newJSONAPIServer initializes and returns a new JSONAPIServer
func newJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

// Run starts the JSONAPIServer
func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

// makeHTTPHandlerFunc is a wrapper that creates a new context and runs the given handler APIFunc
func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

// handleFetchPrice extracts the ticker from URL and fetches the price from the service and returns a JSON response
func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return writeJSON(w, http.StatusOK, &priceResp)
}

// writeJSON returns an encoded JSON given any struct
func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
