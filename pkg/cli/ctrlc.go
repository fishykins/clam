package cli

import (
	"os"
	"os/signal"
	"syscall"
)

func HandleCtrlC(out *chan Message) {
	// Prep ctrl+c handler
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)

	go func() {
		<-s
		// handle app shutdown
		*out <- Message{Sender: "ctrlc"}
	}()
}
