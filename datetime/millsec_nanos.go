package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Today : ", now.Format(time.ANSIC))

	// wrong way to convert nano to millisecond
	nano := now.Nanosecond()
	millisec := nano / 1000000
	fmt.Println("(wrong)Millisecond : ", millisec)

	// correct way to convert time to millisecond - with UnixNano()
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000
	fmt.Println("(correct)Millisecond : ", umillisec)
}
