package connections

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/xrpscan/platform/shutdown"
)

// closeWithTimeout executes a close function with a 3-second timeout
// Also respects the global shutdown context for early termination
func closeWithTimeout(name string, closeFn func() error) {
	// Check if we're already shutting down
	if shutdown.IsActive() {
		shutdownCtx := shutdown.GetContext()
		// Create a timeout context that also respects shutdown context
		ctx, cancel := context.WithTimeout(shutdownCtx, 3*time.Second)
		defer cancel()

		done := make(chan error, 1)
		go func() {
			done <- closeFn()
		}()

		select {
		case err := <-done:
			if err != nil {
				log.Printf("Error closing %s: %v", name, err)
			} else {
				log.Printf("Successfully closed %s", name)
			}
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				log.Printf("Timeout closing %s after 3 seconds", name)
			} else {
				log.Printf("Shutdown cancelled while closing %s", name)
			}
		}
	} else {
		// Fallback to simple timeout if not in shutdown mode
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		done := make(chan error, 1)
		go func() {
			done <- closeFn()
		}()

		select {
		case err := <-done:
			if err != nil {
				log.Printf("Error closing %s: %v", name, err)
			} else {
				log.Printf("Successfully closed %s", name)
			}
		case <-ctx.Done():
			log.Printf("Timeout closing %s after 3 seconds", name)
		}
	}
}

func CloseWriter() {
	closeWithTimeout("Kafka writer", func() error {
		if KafkaWriter != nil {
			return KafkaWriter.Close()
		}
		return nil
	})
}

func CloseReaders() {
	// Define all readers to close with their names
	readers := []struct {
		name   string
		reader interface{ Close() error }
	}{
		{"Kafka Ledger reader", KafkaReaderLedger},
		{"Kafka Transaction reader", KafkaReaderTransaction},
		{"Kafka Validation reader", KafkaReaderValidation},
		{"Kafka PeerStatus reader", KafkaReaderPeerStatus},
		{"Kafka Consensus reader", KafkaReaderConsensus},
		{"Kafka PathFind reader", KafkaReaderPathFind},
		{"Kafka Manifest reader", KafkaReaderManifest},
		{"Kafka Server reader", KafkaReaderServer},
		{"Kafka Default reader", KafkaReaderDefault},
	}

	// Close all readers in parallel with timeout
	var wg sync.WaitGroup
	for _, r := range readers {
		if r.reader != nil {
			wg.Add(1)
			go func(name string, reader interface{ Close() error }) {
				defer wg.Done()
				closeWithTimeout(name, func() error {
					return reader.Close()
				})
			}(r.name, r.reader)
		}
	}

	// Wait for all readers to close or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Overall timeout for all readers
	select {
	case <-done:
		log.Println("All Kafka readers closed")
	case <-time.After(10 * time.Second):
		log.Println("Timeout waiting for Kafka readers to close")
	}
}

func CloseEsClient() {
}

func CloseXrplClient() {
	closeWithTimeout("XRPL client", func() error {
		if XrplClient != nil {
			return XrplClient.Close()
		}
		return nil
	})
}

func CloseAll() {
	log.Println("Closing all connections")
	
	// First unsubscribe from streams
	UnsubscribeStreams()
	
	// Close other connections in parallel
	var wg sync.WaitGroup
	
	// Close writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		CloseWriter()
	}()
	
	// Close readers (already parallel internally)
	wg.Add(1)
	go func() {
		defer wg.Done()
		CloseReaders()
	}()
	
	// Close XRPL client
	wg.Add(1)
	go func() {
		defer wg.Done()
		CloseXrplClient()
	}()
	
	// Close ES client (currently no-op)
	CloseEsClient()
	
	// Wait for all closures with overall timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		log.Println("All connections closed successfully")
	case <-time.After(15 * time.Second):
		log.Println("Timeout waiting for all connections to close")
	}
}
