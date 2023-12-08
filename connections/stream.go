package connections

import (
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
 */
func UnsubscribeStreams() {
	response, err := XrplClient.Unsubscribe([]string{
		xrpl.StreamTypeLedger,
		xrpl.StreamTypeValidations,
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("xrpl.Unsubscribe")
	}
	logger.Log.Debug().Any("status", response["status"]).Any("id", response["id"]).Msg("xrpl.Unsubscribe")
}
