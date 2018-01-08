package bitfinex

import (
	"net/url"
	"strings"
)

// TickersService manages the Tickers endpoint
type TickersService struct {
	client *Client
}

// Get (symbols) - return last Ticker for specified symbols
func (s *TickersService) Get(symbols []string) (TickerSnapshot, error) {
	params := url.Values{}
	params.Set("symbols", strings.Join(symbols, ","))
	req, err := s.client.newRequest("GET", "tickers", params, nil)

	if err != nil {
		return nil, err
	}

	var raw []interface{}
	_, err = s.client.do(req, &raw)
	if err != nil {
		return nil, err
	}

	ts, err := tickerSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
