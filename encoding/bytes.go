package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Metadata struct {
	Metadata map[string]interface{} `json:"Metadata"`
}

func main() {
	jsonMap()
	fmt.Println()
	plainMap()
}

func jsonMap() {
	metadata := Metadata{
		Metadata: make(map[string]interface{}, 3),
	}

	metadata.Metadata["string"] = "this is a string"
	metadata.Metadata["int64"] = int64(42)
	metadata.Metadata["time"] = time.Now()
	fmt.Printf("Metadata: %+v\n", metadata)

	bytes, err := json.Marshal(metadata)
	if err != nil {
		fmt.Printf("Error json marshal: %s\n", err.Error())
		return
	}

	fmt.Printf("JSON bytes: %v\n", bytes)
	fmt.Printf("JSON string: %s\n", string(bytes))
}

func plainMap() {
	values := make(map[string]interface{}, 3)
	values["string"] = "this is a string"
	values["int64"] = int64(42)
	values["time"] = time.Now()
	fmt.Printf("Values: %+v\n", values)

	bytes := []byte(fmt.Sprintf("%+v", values))
	fmt.Printf("Bytes: %v\n", bytes)
}
