package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

const (
	cassandraHost     = "localhost"
	cassandraPort     = 9042
	cassandraKeyspace = "market_data"

	tickerQuery = "SELECT highest_bid,last_price,lowest_ask,event_time,ingestion_time FROM market_data.ticker_updated WHERE source = ? AND exchange = ? AND base_asset_symbol = ? AND base_asset_class = ? AND quote_asset_symbol = ? AND quote_asset_class = ? AND event_time >= ? AND event_time <= ? LIMIT ? ALLOW FILTERING;"
	ohlcvQuery  = "SELECT open,high,low,close,volume,opened_at,closed_at FROM market_data.ohlcv WHERE source = ? AND exchange = ? AND base_asset_symbol = ? AND base_asset_class = ? AND quote_asset_symbol = ? AND quote_asset_class = ? AND interval = ? AND closed_at >= ? AND closed_at <= ? LIMIT ? ALLOW FILTERING;"

	source     = "COINBASE"
	exchange   = "EXCHANGE"
	baseSymbol = "BTC"
	// baseSymbol  = "BAT"
	baseClass   = "CRYPTO_ASSET"
	quoteSymbol = "USDC"
	quoteClass  = "CRYPTO_ASSET"
	interval    = 60
	limit       = 500
)

var session *gocql.Session

func main() {
	createConnection()
	defer session.Close()

	performTickerQuery()
	performOhlcvQuery()
}

func createConnection() {
	clusterCfg := gocql.NewCluster(cassandraHost)
	clusterCfg.Port = cassandraPort
	clusterCfg.Keyspace = cassandraKeyspace
	// clusterCfg.Consistency = gocql.LocalQuorum
	clusterCfg.ConnectTimeout = 5 * time.Second // initial connection timeout, used during initial dial to server (default: 600ms)
	clusterCfg.Timeout = 5 * time.Second        // connection timeout (default: 600ms)
	clusterCfg.NumConns = 5                     // number of connections per host (default: 2)

	var err error
	session, err = clusterCfg.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}

func performTickerQuery() {
	start := time.Date(2020, 4, 29, 3, 20, 0, 0, time.UTC)
	end := time.Date(2020, 4, 29, 3, 30, 0, 0, time.UTC)
	cqlIter := session.Query(
		tickerQuery,
		source, exchange, baseSymbol, baseClass, quoteSymbol, quoteClass,
		start, end, limit).Iter()

	fmt.Printf("Tickers [%s-%s-%s] found: %d\n", source, baseSymbol, quoteSymbol, cqlIter.NumRows())

	closeErr := cqlIter.Close()
	if closeErr != nil {
		fmt.Printf("Error closing ticker connection: %s\n", closeErr.Error())
	}
}

func performOhlcvQuery() {
	start := time.Date(2020, 4, 29, 3, 20, 0, 0, time.UTC)
	end := time.Date(2020, 4, 29, 3, 30, 0, 0, time.UTC)
	cqlIter := session.Query(
		ohlcvQuery,
		source, exchange, baseSymbol, baseClass, quoteSymbol, quoteClass,
		interval, start, end, limit).Iter()

	fmt.Printf("OHLCV [%s-%s-%s] found: %d\n", source, baseSymbol, quoteSymbol, cqlIter.NumRows())

	closeErr := cqlIter.Close()
	if closeErr != nil {
		fmt.Printf("Error closing OHLCV connection: %s\n", closeErr.Error())
	}
}
