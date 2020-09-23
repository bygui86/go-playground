package main

import (
	"fmt"
)

func main() {
	fmtStruct()
	fmtMethodOutput()
}

type Custom struct {
	ID                string          `encoding:"id"`
	Name              string          `encoding:"name"`
	SupportedFeatures map[string]bool `encoding:"supported_features"`
	AvailableFeatures []string        `encoding:"available_features"`
}

func fmtStruct() {
	obj := &Custom{
		ID:                "111",
		Name:              "first",
		SupportedFeatures: map[string]bool{"true": true, "false": false},
	}

	fmt.Printf("object: %+v\n", obj)
	fmt.Printf("available features: %d\n", len(obj.AvailableFeatures))
}

func fmtMethodOutput() {
	first, second, third := outputGenerator()
	fmt.Printf("%s, %s, %d \n", first, second, third)
}

func outputGenerator() (string, string, int) {
	return "first", "second", 3
}
