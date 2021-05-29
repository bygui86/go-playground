package light_listener

import (
	"fmt"
	"os"
)

// OK working
func Listen(syscallCh chan os.Signal, terminationChannel chan bool) {
	fmt.Println("[INFO] Listen for signals: os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL")

	termSign := <-syscallCh
	fmt.Printf("[WARN] Received termination signal: %s \n", termSign.String())
	// <-syscallCh:
	// fmt.Println("[WARN] Received termination signal")
	terminationChannel <- true
}

// KO not working
// func Listen(syscallCh chan os.Signal, terminationChannel chan bool, stopChannel chan bool) {
// 	fmt.Println("[INFO] Listen for signals: os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL")
//
// 	counter := 0 // TODO prevent infinite loop
// 	for {
// 		time.Sleep(1 * time.Second)
// 		// fmt.Println("[DEBUG] Check syscall channel and stop channel")
// 		fmt.Printf("[DEBUG] Check syscall channel and stop channel (counter %d)\n", counter)
//
// 		select {
// 		case termSign := <-syscallCh:
// 			fmt.Printf("[WARN] Received termination signal %s \n", termSign.String())
// 			// case <-syscallCh:
// 			// fmt.Println("[WARN] Received termination signal")
// 			terminationChannel <- true
//
// 		case <-stopChannel:
// 			fmt.Println("[WARN] Received stop signal")
// 			break
//
// 		default:
// 			// fmt.Println("[DEBUG] No syscall or stop signal, continue...")
// 			// continue
// 		}
//
// 		// TODO prevent infinite loop
// 		counter++
// 		if counter > 20 {
// 			fmt.Println("[DEBUG] Counter break, terminating everything!")
// 			terminationChannel <- true
// 			// break
// 		}
// 	}
// }
