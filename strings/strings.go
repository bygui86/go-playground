package main

import (
	"fmt"
	"strings"
)

func main() {
	stringBuilderNewLine()
	stringBuilder()
}

func stringBuilderNewLine() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		b.WriteString(fmt.Sprintf("%d...", i))
	}
	b.WriteString("ignition\n")
	b.WriteString("\n")
	fmt.Println(b.String())
}

func stringBuilder() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
