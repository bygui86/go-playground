package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
	
	"gopkg.in/yaml.v2"
)

const (
	basePath     = "."
	yamlFilePath = "resources/asset-pairs-ranking_v2.yaml"
)

type AssetPairRanking struct {
	Rankings map[string]int `yaml:"assetPairRankings,omitempty"`
}

func main() {
	yamlStart := time.Now()
	yamlRankings := parseYaml()
	fmt.Printf("YAML parsing time: %d nanos \n", time.Now().Sub(yamlStart).Nanoseconds())
	
	fmt.Printf("Rankings: %+v \n", yamlRankings.Rankings)
	fmt.Printf("BTC-USD ranking: %d \n", yamlRankings.Rankings["BTC-USD"])
	fmt.Printf("EOS-USD ranking: %d \n", yamlRankings.Rankings["EOS-USD"])
	
}

func parseYaml() AssetPairRanking {
	path, pathErr := filepath.Rel(basePath, yamlFilePath)
	if pathErr != nil {
		fmt.Printf("loading failed to find folder %s: %s \n", basePath, pathErr)
	}
	
	file, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		fmt.Printf("loading failed to open file %s: %s \n", yamlFilePath, fileErr)
	}
	
	rankings := AssetPairRanking{}
	yamlErr := yaml.Unmarshal(file, &rankings)
	if yamlErr != nil {
		fmt.Printf("loading failed to parse YAML: %s \n", yamlErr)
	}
	return rankings
}
