package termination_listener

import "os"

type Listener struct {
	syscallCh          chan os.Signal
	terminationChannel chan bool
	// stopChannel        chan bool
	active bool
}
