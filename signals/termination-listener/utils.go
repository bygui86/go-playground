package termination_listener

import (
	"fmt"
)

func (l *Listener) listen() {
	fmt.Println("[INFO] Listen for termination signals")

	termSign := <-l.syscallCh
	fmt.Printf("[WARN] Received termination signal: %s \n", termSign.String())
	// <-syscallCh:
	// fmt.Println("[WARN] Received termination signal")
	l.terminationChannel <- true
}
