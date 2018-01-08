package bitfinex

import ()

// TickerService manages the Ticker endpoint
type TickerService struct {
	client *Client
}

// Get (symbol) - return last Ticker for specified symbol
func (s *TickerService) Get(symbol string) (Ticker, error) {
	req, err := s.client.newRequest("GET", "ticker/"+symbol, nil, nil)

	if err != nil {
		return Ticker{}, err
	}

	var raw []interface{}
	_, err = s.client.do(req, &raw)
	if err != nil {
		return Ticker{}, err
	}

	ticker, err := tickerFromRaw(raw)
	if err != nil {
		return Ticker{}, err
	}

	return ticker, nil
}
