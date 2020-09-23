package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	limitKeywordFormat = "LIMIT %d"

	queryTemplate = `SELECT event_time,open,high,low,close,volume,opened_at,closed_at
FROM market_data.ohlcv
WHERE
	source = '%s'
	AND exchange = '%s'
	AND base_asset_symbol = '%s'
	AND base_asset_class = '%s'
	AND quote_asset_symbol = '%s'
	AND quote_asset_class = '%s'
	AND interval = %d
	AND closed_at >= '%s'
	AND closed_at <= '%s'
%s
ALLOW FILTERING`
)

var (
	source           = "COINBASE"
	exchange         = "EXCHANGE"
	baseAssetSymbol  = "BTC"
	baseAssetClass   = "CRYPTO_ASSET"
	quoteAssetSymbol = "USDC"
	quoteAssetClass  = "CRYPTO_ASSET"
	limit            = fmt.Sprintf(limitKeywordFormat, 500)
	interval         = 60
	timeDiff         = 5 * time.Minute
)

func main() {
	startTime := time.Now().UTC().Add(-(timeDiff + 1))
	start := startTime.Format(time.RFC3339)
	end := startTime.Add(-1).Format(time.RFC3339)

	fmt.Println(fmt.Sprintf(
		queryTemplate,
		strings.ToUpper(source), strings.ToUpper(exchange),
		baseAssetSymbol, baseAssetClass,
		quoteAssetSymbol, quoteAssetClass,
		interval, start, end, limit,
	))
}
