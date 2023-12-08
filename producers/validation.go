package producers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

func ProduceValidation(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}

	messageKey := fmt.Sprintf("%s.%s.%s", res["validation_public_key"], res["ledger_index"], res["cookie"])

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicValidations(),
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		logger.Log.Trace().Err(err).Msg("Failed to produce validation")
	}
}
