package producers

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

func ProduceLedger(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}

	messageKey := strconv.Itoa(int(res["ledger_index"].(float64)))

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicLedgers(),
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		logger.Log.Trace().Err(err).Msg("Failed to produce ledger")
	}
}
