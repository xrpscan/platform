package signals

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/shutdown"
)

func HandleAll() {
	// Ensure shutdown context is initialized
	_ = shutdown.GetContext()
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-c
		log.Println("Shutdown signal received, starting graceful shutdown...")
		
		// Cancel the global shutdown context to signal all goroutines
		shutdown.Cancel()
		
		// Give other goroutines a moment to receive the cancellation signal
		time.Sleep(100 * time.Millisecond)
		
		// Create a channel to track completion of cleanup
		cleanupDone := make(chan struct{})
		
		// Start cleanup in a separate goroutine
		go func() {
			defer close(cleanupDone)
			connections.CloseAll()
		}()
		
		// Wait for cleanup to complete or timeout after 30 seconds
		select {
		case <-cleanupDone:
			log.Println("Graceful shutdown completed")
			os.Exit(0)
		case <-time.After(30 * time.Second):
			log.Println("Shutdown timeout exceeded (30s), forcing exit")
			os.Exit(1)
		}
	}()
}
