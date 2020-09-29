package main

import (
	"encoding/json"
	"fmt"
)

const (
	buyPayload = `{
	"type": "trade",
	"timestamp": 1600668910120,
	"price": 42.42,
	"side": "buy"
}`
	sellPayload = `{
	"type": "trade",
	"timestamp": 1600668910120,
	"price": 42.42,
	"side": "sell"
}`

	SideBuy  Side = "buy"
	SideSell Side = "sell"
)

func main() {
	// buyTrade := parsePayloadToTrade([]byte(buyPayload))
	// fmt.Printf("buy trade: %+v \n", *buyTrade)
	// sellTrade := parsePayloadToTrade([]byte(sellPayload))
	// fmt.Printf("sell trade: %+v \n", *sellTrade)

	buyTradeWE := parsePayloadToTradeWithEnum([]byte(buyPayload))
	fmt.Printf("buy trade with enum: %+v \n", *buyTradeWE)
	sellTradeWE := parsePayloadToTradeWithEnum([]byte(sellPayload))
	fmt.Printf("sell trade with enum: %+v \n", *sellTradeWE)
}

func parsePayloadToTrade(payload []byte) *Trade {
	var trade *Trade
	err := json.Unmarshal(payload, &trade)
	if err != nil {
		panic(err)
	}
	return trade
}

func parsePayloadToTradeWithEnum(payload []byte) *TradeWithEnum {
	var trade *TradeWithEnum
	err := json.Unmarshal(payload, &trade)
	if err != nil {
		panic(err)
	}
	return trade
}

type Trade struct {
	Type      string  `json:"type"`
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
	Side      string  `json:"side"`
}

type Side string

type TradeWithEnum struct {
	Type      string  `json:"type"`
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
	Side      Side    `json:"side"`
}
