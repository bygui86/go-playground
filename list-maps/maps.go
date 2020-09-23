package main

import (
	"fmt"
)

func main() {
	// fmt.Println("*** Struct map tests ***")
	// structMap()
	// fmt.Println("")

	// fmt.Println("*** Bool map tests ***")
	// boolMap()
	// fmt.Println("")

	fmt.Println("*** Int map tests ***")
	intMap()
	fmt.Println("")
}

func intMap() {
	intMap := map[string]int{"one": 1, "two": 2, "three": 3}

	position1 := intMap["one"]
	fmt.Printf("value for one: [%d] \n", position1) // expected [1]
	position2, positionFound2 := intMap["two"]
	fmt.Printf("value for two: [%d][%t] \n", position2, positionFound2) // expected [2][TRUE]
	position4, positionFound4 := intMap["four"]
	fmt.Printf("value for two: [%d][%t] \n", position4, positionFound4) // expected [?][FALSE]
}

func boolMap() {
	var m = make(map[string]bool)

	m["one"] = true
	m["two"] = false

	if m == nil {
		fmt.Println("m is nil")
		return
	}
	fmt.Println(m)

	value1, found1 := m["one"]
	fmt.Println("value for one:", value1) // expected TRUE
	fmt.Println("found for one:", found1) // expected TRUE
	value2, found2 := m["two"]
	fmt.Println("value for two:", value2) // expected FALSE
	fmt.Println("found for two:", found2) // expected TRUE
	value3, found3 := m["three"]
	fmt.Println("value for three:", value3) // expected FALSE
	fmt.Println("found for three:", found3) // expected FALSE

	valueOnly1 := m["one"]
	fmt.Println("value ONLY for one:", valueOnly1) // expected TRUE
	valueOnly3 := m["three"]
	fmt.Println("value ONLY for three:", valueOnly3) // expected FALSE
	_, foundOnly2 := m["two"]
	fmt.Println("found ONLY for two:", foundOnly2) // expected TRUE
	_, foundOnly4 := m["four"]
	fmt.Println("found ONLY for four:", foundOnly4) // expected FALSE
}

type Custom struct {
	ID                string          `encoding:"id"`
	Name              string          `encoding:"name"`
	SupportedFeatures map[string]bool `encoding:"supported_features"`
	AvailableFeatures []string        `encoding:"available_features"`
}

func structMap() {
	var m = make(map[string]*Custom)
	m["111"] = &Custom{
		ID:   "111",
		Name: "first",
	}
	m["222"] = &Custom{
		ID:   "222",
		Name: "second",
	}

	if m == nil {
		fmt.Println("m is nil")
		return
	}
	fmt.Println(m)

	value1, found1 := m["111"]
	fmt.Printf("value for 111: %+v\n", *value1) // expected 111:first
	fmt.Println("found for 111:", found1)       // expected TRUE
	value2, found2 := m["222"]
	fmt.Printf("value for 222: %+v\n", *value2) // expected 222:second
	fmt.Println("found for 222:", found2)       // expected TRUE
	value3, found3 := m["333"]
	if value3 != nil {
		fmt.Printf("value for 333: %+v\n", *value3)
	} else {
		fmt.Println("value for 333: NIL") // expected
	}
	fmt.Println("found for 333:", found3) // expected FALSE
}
