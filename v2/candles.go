package bitfinex

import (
	"net/url"
	"strconv"
)

// CandlesService manages the candles endpoint
type CandlesService struct {
	client *Client
}

// Hist (timeFrame, symbol, limit, start, end, sort) - return hist candles
func (s *CandlesService) Hist(timeFrame string, symbol string, limit int, start uint64, end uint64, sort int) (CandleSnapshot, error) {
	params := url.Values{}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if start > 0 {
		params.Set("start", strconv.FormatUint(start, 10))
	}
	if end > 0 {
		params.Set("end", strconv.FormatUint(end, 10))
	}
	params.Set("sort", strconv.Itoa(sort))

	url := "candles/trade:" + timeFrame + ":" + symbol + "/hist"
	req, err := s.client.newRequest("GET", url, params, nil)

	if err != nil {
		return nil, err
	}

	var raw []interface{}
	_, err = s.client.do(req, &raw)
	if err != nil {
		return nil, err
	}

	ts, err := candleSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
