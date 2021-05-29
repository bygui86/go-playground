package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

const filePath = "clients.csv"

var data []*Client

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id        string             `csv:"client_id"`
	Name      string             `csv:"client_name"`
	Age       int64              `csv:"client_age"`
	NotUsed   string             `csv:"-"`
	Addresses []*Address         `csv:"-"`
	Mapped    map[string]*Mapped `csv:"-"`
}

type Address struct {
	City    string `csv:"city"`
	Country string `csv:"country"`
}

type Mapped struct {
	Line  string  `csv:"line"`
	Value float64 `csv:"value"`
}

func main() {

	createData()

	writeDataToCsvString()

	writeDataToCsvFile()

	readDataFromCsvFile()

	// if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
	// 	panic(err)
	// }
}

func createData() {
	addresses := make([]*Address, 3)
	addresses[0] = &Address{City: "Milan", Country: "IT"}
	addresses[1] = &Address{City: "Zurich", Country: "CH"}
	addresses[2] = &Address{City: "New York", Country: "USA"}

	mapped := make(map[string]*Mapped, 3)
	mapped["A"] = &Mapped{Line: "line1", Value: 43.333}
	mapped["B"] = &Mapped{Line: "line2", Value: 42.22}
	mapped["C"] = &Mapped{Line: "line3", Value: 31.1}

	data = append(data, &Client{Id: "12", Name: "John", Age: 21, Addresses: addresses, Mapped: mapped})
	data = append(data, &Client{Id: "13", Name: "Fred", Age: 35, Addresses: addresses})
	data = append(data, &Client{Id: "15", Name: "Danny", Age: 42})
}

func writeDataToCsvString() {
	csvContent, err := gocsv.MarshalString(&data) // Get all data as CSV string
	if err != nil {
		panic(err)
	}

	log.Println(csvContent) // Display all data as CSV string
}

func writeDataToCsvFile() {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gocsv.MarshalFile(&data, file) // Get all data as CSV string and save to a file
	if err != nil {
		panic(err)
	}
}

func readDataFromCsvFile() {

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &data); err != nil { // Load data from file
		panic(err)
	}

	for i, client := range data {
		addresses := strings.Builder{}
		addresses.WriteString("[ ")
		for _, address := range client.Addresses {
			addresses.WriteString(fmt.Sprintf("City=%s Country=%s ", address.City, address.Country))
		}
		addresses.WriteString("]")

		mapped := strings.Builder{}
		mapped.WriteString("[ ")
		for key, value := range client.Mapped {
			mapped.WriteString(fmt.Sprintf("Key=%s Line=%s Value=%.3f ", key, value.Line, value.Value))
		}
		mapped.WriteString("]")

		fmt.Printf("Client %d: ID=%s Name=%s Age=%d Addresses=%s Mapped=%s \n",
			i, client.Id, client.Name, client.Age, addresses.String(), mapped.String())
	}
}
