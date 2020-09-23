package main

import "fmt"

func main() {
	noMismatch()
}

func noMismatch() {
	// best case: no mismatch between sourceList and rankingMap
	sourceList := buildSourceList()
	rankingMap := buildRankingMap()
	resultList := make([]string, len(sourceList))

	tail := len(sourceList) - 1
	for _, name := range sourceList {
		position := rankingMap[name]
		if position != 0 {
			resultList[position-1] = name
		} else {
			resultList[tail] = name
			tail--
		}
	}

	fmt.Printf("Result list: %v", resultList)
}

func buildSourceList() []string {
	return []string{"XRP-USD", "LTC-BTC", "XLM-USD", "XRP-BTC", "BTC-USDC", "ETH-BTC", "ETH-USD", "BTC-USD", "ETH-USDC", "LTC-USD"}
}

func buildSourceListWithAdditional() []string {
	return []string{"XRP-USD", "LTC-BTC", "XLM-USD", "XRP-BTC", "BTC-USDC", "ETH-BTC", "ETH-USD", "BTC-USD", "ETH-USDC", "LTC-USD", "BYG-UIN"}
}

func buildRankingMap() map[string]int {
	return map[string]int{"BTC-USD": 1, "BTC-USDC": 2, "ETH-BTC": 3, "ETH-USD": 4, "ETH-USDC": 5, "LTC-BTC": 6, "LTC-USD": 7, "XRP-BTC": 8, "XRP-USD": 9, "XLM-USD": 10}
}
