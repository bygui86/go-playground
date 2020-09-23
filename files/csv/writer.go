package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	filePath = "./result.csv"
)

var (
	headers = []string{"EventTime", "Open", "High", "Low", "Close", "Volume"}
)

type Ohlcv struct {
	EventTime time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

func (o *Ohlcv) GetCsvValues() []string {
	values := make([]string, 6)
	values[0] = o.EventTime.String()
	values[1] = fmt.Sprintf("%.8f", o.Open)
	values[2] = fmt.Sprintf("%.8f", o.High)
	values[3] = fmt.Sprintf("%.8f", o.Low)
	values[4] = fmt.Sprintf("%.8f", o.Close)
	values[5] = fmt.Sprintf("%.8f", o.Volume)
	return values
}

func main() {
	log.Println("CSV writer started")

	log.Println("Create CSV file")
	file, fErr := os.Create(filePath)
	if fErr != nil {
		log.Fatal("Cannot create file", fErr)
	}
	defer file.Close()

	log.Println("Open CSV file writer")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	wErr := writer.Write(headers)
	if wErr != nil {
		log.Fatal("Cannot write to file", fErr)
	}

	for i := 0; i < 10000; i++ {
		// err := writer.Write([]string{fmt.Sprintf("hello %d", i)})
		err := writer.Write(generateOhlcv().GetCsvValues())
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
		if i%10 == 0 {
			log.Printf("Wrote %d lines to CSV file\n", i)
		}
		writer.Flush()
		time.Sleep(3 * time.Second)
	}

	log.Println("CSV writer completed")
}

func generateOhlcv() *Ohlcv {
	return &Ohlcv{
		EventTime: time.Now(),
		Open:      1.42,
		High:      2.42,
		Low:       3.42,
		Close:     4.42,
		Volume:    5.42,
	}
}
