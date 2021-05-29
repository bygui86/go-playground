package termination_listener

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(terminationChannel chan bool) *Listener {
	fmt.Println("[INFO] Create new termination signals listener")

	return &Listener{
		syscallCh:          make(chan os.Signal),
		terminationChannel: terminationChannel,
		// stopChannel:        make(chan bool, 1),
	}
}

func (l *Listener) Start() {
	if !l.active {
		fmt.Println("[INFO] Start termination signals listener (os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)")

		signal.Notify(l.syscallCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

		go l.listen()

		l.active = true
		fmt.Printf("[DEBUG] Termination signals listener active %t \n", l.active)
	} else {
		fmt.Println("[WARN] Termination signals listener already running")
	}
}

func (l *Listener) Shutdown(timeout time.Duration) {
	if l.active {
		fmt.Printf("[WARN] Shutdown termination signals listener, timeout %.0f sec \n", timeout.Seconds())

		// l.stopChannel <- true

		time.Sleep(timeout)

		// close(l.stopChannel)
		close(l.syscallCh)
	} else {
		fmt.Println("[WARN] Termination signals listener not running yet")
	}
}
