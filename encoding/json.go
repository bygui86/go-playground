package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	tickerJsonBody = `{
	"stream": "btcusdt@ticker",
	"data": {
		"e": "24hrTicker",
		"E": 1588770693638,
		"s": "BTCUSDT",
		"p": "366.50000000",
		"P": "4.126",
		"w": "9072.90323879",
		"x": "8881.39000000",
		"c": "9248.15000000",
		"Q": "0.00000700",
		"b": "9247.56000000",
		"B": "0.17283700",
		"a": "9248.15000000",
		"A": "0.71341400",
		"o": "8881.65000000",
		"h": "9380.00000000",
		"l": "8811.00000000",
		"v": "89530.25332400",
		"q": "812299325.35257007",
		"O": 1588684293342,
		"C": 1588770693342,
		"F": 309710193,
		"L": 310618203,
		"n": 908011
	}
}
`
	ohlcvJsonBody = `{
	"stream": "btcusdt@kline_1m",
	"data": {
		"e": "kline",
		"E": 1588785457641,
		"s": "BTCUSDT",
		"k": {
			"t": 1588785420000,
			"T": 1588785479999,
			"s": "BTCUSDT",
			"i": "1m",
			"f": 310793488,
			"L": 310793761,
			"o": "9241.94000000",
			"c": "9238.16000000",
			"h": "9241.99000000",
			"l": "9236.02000000",
			"v": "20.75690000",
			"n": 274,
			"x": false,
			"q": "191791.60679415",
			"V": "5.12292400",
			"Q": "47331.61704114",
			"B": "0"
		}
	}
}
`

	tickerJsonBodyIncomplete = `{
	"stream": "btcusdt@ticker"
}`
	ohlcvJsonBodyIncomplete = `{
	"stream": "btcusdt@kline_1m",
	"data": {
		"e": "kline",
		"E": 1588785457641,
		"s": "BTCUSDT"
	}
}`
)

func main() {
	// ticker()
	ohlcv()
}

func ticker() {
	start := time.Now()
	var ticker BinanceStreamTicker
	err := json.Unmarshal([]byte(tickerJsonBody), &ticker)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	diff := time.Since(start)
	fmt.Printf("Time elapsed: sec [%f] mil [%d] nanos [%d]\n", diff.Seconds(), diff.Milliseconds(), diff.Nanoseconds())

	fmt.Printf("Ticker: %+v\n", ticker)
	fmt.Printf("Ticker float64: %f\n", ticker.Data.PriceChange)
	fmt.Printf("Ticker int64: %d\n", ticker.Data.EventTime)
	eventTimeMil := time.Unix(0, ticker.Data.EventTime*int64(time.Millisecond))
	fmt.Printf("Ticker event time: %+v\n", eventTimeMil)
}

type BinanceStreamTicker struct {
	Stream string        `encoding:"stream"`
	Data   BinanceTicker `encoding:"data"`
}

type BinanceTicker struct {
	EventType                   string  `encoding:"e"`
	EventTime                   int64   `encoding:"E"` // unix milliseconds
	Symbol                      string  `encoding:"s"`
	PriceChange                 float64 `encoding:"p,string"`
	PriceChangePercent          float64 `encoding:"P,string"`
	WeightedAveragePrice        float64 `encoding:"w,string"`
	FirstTrade                  float64 `encoding:"x,string"`
	LastPrice                   float64 `encoding:"c,string"`
	LastQuantity                float64 `encoding:"Q,string"`
	BestBidPrice                float64 `encoding:"b,string"`
	BestBidQuantity             float64 `encoding:"B,string"`
	BestAskPrice                float64 `encoding:"a,string"`
	BestAskQuantity             float64 `encoding:"A,string"`
	OpenPrice                   float64 `encoding:"o,string"`
	HighPrice                   float64 `encoding:"h,string"`
	LowPrice                    float64 `encoding:"l,string"`
	TotalTradedBaseAssetVolume  float64 `encoding:"v,string"`
	TotalTradedQuoteAssetVolume float64 `encoding:"q,string"`
	StatisticsOpenTime          int64   `encoding:"O"`
	StatisticsCloseTime         int64   `encoding:"C"`
	FirstTradeID                int64   `encoding:"F"`
	LastTradeID                 int64   `encoding:"L"`
	TradesTotalNumber           int64   `encoding:"n"`
}

func ohlcv() {
	start := time.Now()
	var ohlcv BinanceStreamKline
	err := json.Unmarshal([]byte(ohlcvJsonBody), &ohlcv)
	// err := encoding.Unmarshal([]byte(ohlcvJsonBodyIncomplete), &ohlcv)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	diff := time.Since(start)
	fmt.Printf("Time elapsed: sec [%f] mil [%d] nanos [%d]\n", diff.Seconds(), diff.Milliseconds(), diff.Nanoseconds())

	fmt.Printf("OHLCV: %+v\n", ohlcv)

	fmt.Printf("OHLCV float64: %f\n", ohlcv.Data.Sample.Open)
	fmt.Printf("OHLCV int64: %d\n", ohlcv.Data.EventTime)
	fmt.Printf("OHLCV int64: %d\n", ohlcv.Data.Sample.ClosedAt)
	eventTimeMil := time.Unix(0, ohlcv.Data.EventTime*int64(time.Millisecond))
	closedAtMil := time.Unix(0, ohlcv.Data.Sample.ClosedAt*int64(time.Millisecond))
	fmt.Printf("OHLCV event time: %+v\n", eventTimeMil)
	fmt.Printf("OHLCV closed at: %+v\n", closedAtMil)
}

type BinanceStreamKline struct {
	Stream string       `encoding:"stream"`
	Data   BinanceKline `encoding:"data"`
}
type BinanceKline struct {
	EventType string       `encoding:"e"`
	EventTime int64        `encoding:"E"` // unix milliseconds
	Symbol    string       `encoding:"s"`
	Sample    BinanceOhlcv `encoding:"k"`
}
type BinanceOhlcv struct {
	OpenedAt                 int64   `encoding:"t"` // unix milliseconds
	ClosedAt                 int64   `encoding:"T"` // unix milliseconds
	Symbol                   string  `encoding:"s"`
	Interval                 string  `encoding:"i"`
	FirstTradeID             int64   `encoding:"f"`
	LastTradeID              int64   `encoding:"L"`
	Open                     float64 `encoding:"o,string"`
	Close                    float64 `encoding:"c,string"`
	High                     float64 `encoding:"h,string"`
	Low                      float64 `encoding:"l,string"`
	Volume                   float64 `encoding:"v,string"`
	NumberOfTrades           int64   `encoding:"n"`
	SampleClosed             bool    `encoding:"x"`
	QuoteAssetVolume         float64 `encoding:"q,string"`
	TakerBuyBaseAssetBolume  float64 `encoding:"V,string"`
	TakerBuyQuoteAssetVolume float64 `encoding:"Q,string"`
}
