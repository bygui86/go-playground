package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	// stringElementNumber = 58
	// intElementNumber    = 10

	stringElementNumber = 10
	intElementNumber    = 5
)

var (
	stringKeys  []string
	multidimMap map[string]map[int64]bool
)

func main() {
	fmt.Println()
	fmt.Println()

	prepareMap()
	fmt.Printf("Map %+v \n", multidimMap)

	rightBound := time.Now()
	key := "CVC-USDC"
	updateKeyValues(key, buildSamples(&rightBound, false), &rightBound)
	// fmt.Printf("Map %s has %d elements: %+v \n", key, len(multidimMap[key]), multidimMap[key])
	printKeyMap(key)

	newRightBound := rightBound.Add(1 * time.Minute)
	updateKeyValues(key, buildSamples(&newRightBound, true), &newRightBound)
	// fmt.Printf("Map %s has %d elements: %+v \n", key, len(multidimMap[key]), multidimMap[key])
	printKeyMap(key)

	fmt.Println()
}

func prepareMap() {
	stringKeys = buildReducedStringKeys()
	multidimMap = make(map[string]map[int64]bool, stringElementNumber)
	for _, key := range stringKeys {
		multidimMap[key] = make(map[int64]bool, intElementNumber)
	}
}

func updateKeyValues(key string, samples map[int64]bool, sentinel *time.Time) {
	delete(multidimMap[key], sentinel.Add(-time.Duration(intElementNumber+1)*time.Minute).Unix())
	for newSampleKey, newSampleValue := range samples {
		oldSampleValue, oldSampleValueFound := multidimMap[key][newSampleKey]
		if !oldSampleValueFound || (!oldSampleValue && newSampleValue) {
			multidimMap[key][newSampleKey] = newSampleValue
		}
	}
}

func updateSingleValue(assetPair string, time int64, published bool) {
	if !multidimMap[assetPair][time] {
		multidimMap[assetPair][time] = published
	}
}

func buildSamples(rightBound *time.Time, random bool) map[int64]bool {
	ohlcvSamples := make(map[int64]bool, intElementNumber)
	for i := 1; i < intElementNumber+1; i++ {
		value := false
		if random {
			if i%3 == 0 {
				value = true
			}
		}
		ohlcvSamples[rightBound.Add(-time.Duration(i)*time.Minute).Unix()] = value
	}
	return ohlcvSamples
}

func printKeyMap(key string) {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("\n")
	for timeKey := range multidimMap[key] {

		strBuilder.WriteString(
			fmt.Sprintf("\t %+v:%t \n", time.Unix(timeKey, 0), multidimMap[key][timeKey]),
		)
	}
	strBuilder.WriteString("\n")
	fmt.Printf("Map %s has %d elements:%s\n", key, len(multidimMap[key]), strBuilder.String())
}

func buildReducedStringKeys() []string {
	return []string{
		"EOS-USD",
		"DNT-USDC",
		"KNC-USD",
		"EOS-EUR",
		"LOOM-USDC",
		"LINK-USD",
		"CVC-USDC",
		"EOS-BTC",
		"XTZ-USD",
		"ETC-GBP",
	}
}

func buildExtendedStringKeys() []string {
	return []string{
		"EOS-USD",
		"DNT-USDC",
		"KNC-USD",
		"EOS-EUR",
		"LOOM-USDC",
		"LINK-USD",
		"CVC-USDC",
		"EOS-BTC",
		"XTZ-USD",
		"ETC-GBP",
		"MANA-USDC",
		"GNT-USDC",
		"ZEC-USDC",
		"XLM-USD",
		"ETH-BTC",
		"LINK-ETH",
		"BCH-EUR",
		"BCH-GBP",
		"KNC-BTC",
		"ETH-GBP",
		"XRP-EUR",
		"ETH-EUR",
		"BTC-USDC",
		"ZRX-BTC",
		"ZRX-USD",
		"BCH-USD",
		"ETH-USD",
		"ATOM-USD",
		"LTC-BTC",
		"OXT-USD",
		"XLM-EUR",
		"ALGO-USD",
		"XLM-BTC",
		"LTC-USD",
		"DASH-BTC",
		"ETC-BTC",
		"LTC-EUR",
		"BTC-EUR",
		"DAI-USDC",
		"ETH-DAI",
		"XTZ-BTC",
		"BCH-BTC",
		"BTC-USD",
		"LTC-GBP",
		"BAT-ETH",
		"REP-BTC",
		"BAT-USDC",
		"ETC-USD",
		"ZEC-BTC",
		"BTC-GBP",
		"XRP-BTC",
		"ETC-EUR",
		"ETH-USDC",
		"ATOM-BTC",
		"REP-USD",
		"DASH-USD",
		"ZRX-EUR",
		"XRP-USD",
	}
}
