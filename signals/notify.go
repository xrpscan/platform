package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xrpscan/platform/connections"
)

func HandleAll() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		connections.CloseAll()
		os.Exit(0)
	}()
}
