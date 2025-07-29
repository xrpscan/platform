package connections

import (
	"context"
	"time"

	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

/*
* TLDR - Do not subscribe to xrpl.StreamTypeTransactions
*
* XRPL `transaction` stream messages are incompatible with rippled's native
* transaction format. Therefore, this service does not process transactions
* streamed on `xrpl.StreamTypeTransactions` stream. Instead, we listen to the
* ledger stream, fetch transactions from rippled, and add those transactions to
* the Kafka topic.
 */
func SubscribeStreams() {
	response, err := XrplClient.Subscribe([]string{
		xrpl.StreamTypeLedger,
		xrpl.StreamTypeValidations,
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("xrpl.Subscribe")
	}
	if response["status"].(string) == "error" {
		logger.Log.Error().Any("error", response["error"]).Any("id", response["id"]).Any("error_message", response["error_message"]).Msg("xrpl.Subscribe")
	} else {
		logger.Log.Debug().Any("status", response["status"]).Any("id", response["id"]).Any("result", response["result"]).Msg("xrpl.Subscribe")
	}
}

/*
* Unsubscribe XRPL streams (usually before disconnecting)
* Uses timeout to prevent hanging during shutdown
 */
func UnsubscribeStreams() {
	// Skip unsubscribe if client is nil
	if XrplClient == nil {
		logger.Log.Debug().Msg("xrpl.Unsubscribe skipped - client is nil")
		return
	}

	// Create a timeout context for the unsubscribe operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Channel to handle the unsubscribe result
	done := make(chan struct{})
	var response map[string]interface{}
	var err error

	// Run unsubscribe in a goroutine with timeout
	go func() {
		defer close(done)
		response, err = XrplClient.Unsubscribe([]string{
			xrpl.StreamTypeLedger,
			xrpl.StreamTypeValidations,
		})
	}()

	// Wait for completion or timeout
	select {
	case <-done:
		if err != nil {
			logger.Log.Error().Err(err).Msg("xrpl.Unsubscribe")
		} else {
			logger.Log.Debug().Any("status", response["status"]).Any("id", response["id"]).Msg("xrpl.Unsubscribe")
		}
	case <-ctx.Done():
		logger.Log.Warn().Msg("xrpl.Unsubscribe timed out after 5 seconds")
	}
}
