package main

import "fmt"

const (
	nilInput   Input = ""
	emptyInput Input = "empty"
	dogInput   Input = "dog"
	catInput   Input = "cat"
	batInput   Input = "bat"
)

type Input string

func main() {

	input := nilInput
	// input := emptyInput
	// input := dogInput
	// input := catInput
	// input := batInput

	switch input {
	case emptyInput:
		break
	case dogInput, catInput:
		fmt.Println("i'm a pet")
	case batInput:
		fmt.Println("i'm batman!")
	default:
		fmt.Println("default :(")
	}
}
