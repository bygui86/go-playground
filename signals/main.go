package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	lightListener "github.com/bygui86/go-playground/signals/light-listener"
	terminationListener "github.com/bygui86/go-playground/signals/termination-listener"
)

func main() {
	// simpleExample() // OK

	// advancedPlainExample() // OK

	advancedStructExample() // OK
}

// ADVANCED

func advancedStructExample() {
	fmt.Println("[INFO] Start advanced STRUCT example")

	terminationChannel := make(chan bool, 1)

	termSigListener := terminationListener.New(terminationChannel)

	termSigListener.Start()

	// either
	// fmt.Println("[INFO] Waiting for termination signal")
	// <-terminationChannel

	// or
	time.Sleep(time.Duration(5) * time.Second)
	termSigListener.Shutdown(time.Duration(5) * time.Second)

	fmt.Println("[INFO] Shutdown advanced STRUCT example")
}

func advancedPlainExample() {
	fmt.Println("[INFO] Start advanced PLAIN example")

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	terminationChannel := make(chan bool, 1)
	// stopChannel := make(chan bool, 1)

	// go lightListener.Listen(syscallCh, terminationChannel, stopChannel)
	go lightListener.Listen(syscallCh, terminationChannel)

	fmt.Println("[INFO] Waiting for termination signal")
	<-terminationChannel

	fmt.Println("[INFO] Shutdown advanced PLAIN example")
	// time.Sleep(time.Duration(5) * time.Second)
	// stopChannel <- true
}

// SIMPLE

func simpleExample() {
	// Setup syscall simpleListener
	setupSimpleListener()

	// Run our program
	for {
		fmt.Println("Run some job")
		time.Sleep(10 * time.Second)
	}
}

// setupSimpleListener creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupSimpleListener() {
	syscallCh := make(chan os.Signal)
	// signal.Notify(syscallCh, os.Interrupt, syscall.SIGTERM)
	signal.Notify(syscallCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	go simpleListener(syscallCh)
}

func simpleListener(syscallCh chan os.Signal) {
	<-syscallCh
	fmt.Println("\r   Ctrl+C pressed in Terminal")
	os.Exit(0)
}
