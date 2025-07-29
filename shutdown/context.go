package shutdown

import (
	"context"
	"sync"
)

var (
	// Global shutdown context and cancel function
	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc
	shutdownOnce   sync.Once
)

// GetContext returns the global shutdown context
// This context is cancelled when a shutdown signal is received
func GetContext() context.Context {
	shutdownOnce.Do(func() {
		shutdownCtx, shutdownCancel = context.WithCancel(context.Background())
	})
	return shutdownCtx
}

// Cancel cancels the global shutdown context
// This should only be called by the signal handler
func Cancel() {
	shutdownOnce.Do(func() {
		shutdownCtx, shutdownCancel = context.WithCancel(context.Background())
	})
	if shutdownCancel != nil {
		shutdownCancel()
	}
}

// IsActive returns true if shutdown has been initiated
func IsActive() bool {
	ctx := GetContext()
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// WaitForShutdown blocks until a shutdown signal is received
// This can be used by the main function to keep the program running
func WaitForShutdown() {
	ctx := GetContext()
	<-ctx.Done()
}