package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// https://golangcode.com/how-to-read-a-csv-file-into-a-struct/

const (
	basePath    = "."
	csvFilePath = "resources/asset-pairs_ranking.csv-raw"
)

type AssetPairRanking struct {
	Rankings map[string]int `yaml:"assetPairRankings,omitempty"`
}

func main() {
	csvStart := time.Now()
	csvRankings := parseCsv()
	fmt.Printf("CSV parsing time: %d nanos \n", time.Now().Sub(csvStart).Nanoseconds())

	fmt.Printf("[CSV] Rankings: %+v \n", csvRankings.Rankings)
	fmt.Printf("[CSV] BTC-USD ranking: %d \n", csvRankings.Rankings["BTC-USD"])
	fmt.Printf("[CSV] EOS-USD ranking: %d \n", csvRankings.Rankings["EOS-USD"])
}

func parseCsv() AssetPairRanking {
	file, openErr := os.Open(csvFilePath)
	if openErr != nil {
		fmt.Printf("loading failed to open file %s: %s \n", csvFilePath, openErr)
	}
	defer file.Close()

	lines, readErr := csv.NewReader(file).ReadAll()
	if readErr != nil {
		fmt.Printf("loading failed to read file %s: %s \n", csvFilePath, readErr)
	}

	rankings := AssetPairRanking{
		Rankings: make(map[string]int, len(lines)),
	}
	for _, line := range lines {
		value, _ := strconv.Atoi(line[2])
		rankings.Rankings[line[0]] = value
	}
	return rankings
}
