package bitfinex

import (
	"fmt"
)

type Ticker struct {
	Symbol          string
	FRR             float64
	Bid             float64
	BidPeriod       int64
	BidSize         float64
	Ask             float64
	AskPeriod       int64
	AskSize         float64
	DailyChange     float64
	DailyChangePerc float64
	LastPrice       float64
	Volume          float64
	High            float64
	Low             float64
}

func tickerFromRaw(raw []interface{}) (Ticker, error) {
	var ticker Ticker
	switch len(raw) {
	case 10:
		ticker = Ticker{
			Bid:             f64ValOrZero(raw[0]),
			BidSize:         f64ValOrZero(raw[1]),
			Ask:             f64ValOrZero(raw[2]),
			AskSize:         f64ValOrZero(raw[3]),
			DailyChange:     f64ValOrZero(raw[4]),
			DailyChangePerc: f64ValOrZero(raw[5]),
			LastPrice:       f64ValOrZero(raw[6]),
			Volume:          f64ValOrZero(raw[7]),
			High:            f64ValOrZero(raw[8]),
			Low:             f64ValOrZero(raw[9]),
		}
	case 11:
		ticker = Ticker{
			Symbol:          sValOrEmpty(raw[0]),
			Bid:             f64ValOrZero(raw[1]),
			BidSize:         f64ValOrZero(raw[2]),
			Ask:             f64ValOrZero(raw[3]),
			AskSize:         f64ValOrZero(raw[4]),
			DailyChange:     f64ValOrZero(raw[5]),
			DailyChangePerc: f64ValOrZero(raw[6]),
			LastPrice:       f64ValOrZero(raw[7]),
			Volume:          f64ValOrZero(raw[8]),
			High:            f64ValOrZero(raw[9]),
			Low:             f64ValOrZero(raw[10]),
		}
	case 13:
		ticker = Ticker{
			FRR:             f64ValOrZero(raw[0]),
			Bid:             f64ValOrZero(raw[1]),
			BidSize:         f64ValOrZero(raw[2]),
			BidPeriod:       i64ValOrZero(raw[3]),
			Ask:             f64ValOrZero(raw[4]),
			AskSize:         f64ValOrZero(raw[5]),
			AskPeriod:       i64ValOrZero(raw[6]),
			DailyChange:     f64ValOrZero(raw[7]),
			DailyChangePerc: f64ValOrZero(raw[8]),
			LastPrice:       f64ValOrZero(raw[9]),
			Volume:          f64ValOrZero(raw[10]),
			High:            f64ValOrZero(raw[11]),
			Low:             f64ValOrZero(raw[12]),
		}
	case 14:
		ticker = Ticker{
			Symbol:          sValOrEmpty(raw[0]),
			FRR:             f64ValOrZero(raw[1]),
			Bid:             f64ValOrZero(raw[2]),
			BidSize:         f64ValOrZero(raw[3]),
			BidPeriod:       i64ValOrZero(raw[4]),
			Ask:             f64ValOrZero(raw[5]),
			AskSize:         f64ValOrZero(raw[6]),
			AskPeriod:       i64ValOrZero(raw[7]),
			DailyChange:     f64ValOrZero(raw[8]),
			DailyChangePerc: f64ValOrZero(raw[9]),
			LastPrice:       f64ValOrZero(raw[10]),
			Volume:          f64ValOrZero(raw[11]),
			High:            f64ValOrZero(raw[12]),
			Low:             f64ValOrZero(raw[13]),
		}
	default:
		return Ticker{}, fmt.Errorf("data slice is not compatible with ticker: %#v", raw)
	}

	return ticker, nil
}

type TickerUpdate Ticker
type TickerSnapshot []Ticker

func tickerSnapshotFromRaw(raw []interface{}) (ts TickerSnapshot, err error) {
	if len(raw) == 0 {
		return
	}

	switch raw[0].(type) {
	case []interface{}:
		for _, v := range raw {
			if l, ok := v.([]interface{}); ok {
				t, err := tickerFromRaw(l)
				if err != nil {
					return ts, err
				}
				ts = append(ts, t)
			}
		}
	default:
		return ts, fmt.Errorf("not a ticker snapshot")
	}

	return
}

//type Trade struct {
//ID     int64
//MTS    int64
//Amount float64
//Price  float64
//Rate   float64
//Period int64
//}
