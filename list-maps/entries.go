package main

import (
	"fmt"
	"strconv"
)

func main() {
	entries := buildList()
	printList(entries)
	fmt.Println()
	pointerEntries := buildPointerList()
	printPointerList(pointerEntries)
}

type Entry struct {
	id            int
	name          string
	age           int
	registered    bool
	totalPurchase float64
}

func buildList() []Entry {
	list := make([]Entry, 10)
	var entry Entry
	for i := 1; i <= 10; i++ {
		entry.id = i
		entry.name = "account " + strconv.Itoa(i)
		entry.age = i
		if i%2 == 0 {
			entry.registered = true
		} else {
			entry.registered = false
		}
		entry.totalPurchase = float64(i)
		list[i-1] = entry
	}
	return list
}

func printList(list []Entry) {
	fmt.Printf("Entries: %+v \n", list)
}

func buildPointerList() []*Entry {
	list := make([]*Entry, 10)
	var entry Entry
	for i := 1; i <= 10; i++ {
		entry.id = i
		entry.name = "account " + strconv.Itoa(i)
		entry.age = i
		if i%2 == 0 {
			entry.registered = true
		} else {
			entry.registered = false
		}
		entry.totalPurchase = float64(i)
		list[i-1] = &entry
	}
	return list
}

func printPointerList(list []*Entry) {
	fmt.Println("Entries (pointers):")
	for i := range list {
		fmt.Printf("%+v \n", list[i])
	}
}
