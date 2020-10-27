package dumper

import (
	"log"
)

// Config -
type Config struct {
	Topics   []string
	FileName string
	Proto    bool
	JSON     bool
}

// NewConfig - Setup a new Dumper configuration
func NewConfig(topics []string, fileName string, proto bool, json bool) *Config {

	log.Println("Setup new Dumper config...")

	return &Config{
		Topics:   topics,
		FileName: fileName,
		Proto:    proto,
		JSON:     json,
	}
}
