package producers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

func ProduceTx(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}
	ledgerIndex := strconv.Itoa(int(res["ledger_index"].(float64)))
	requestId := fmt.Sprintf("ledger.%s.tx", ledgerIndex)

	request := xrpl.BaseRequest{
		"id":           requestId,
		"command":      "ledger",
		"ledger_index": ledgerIndex,
		"transactions": true,
		"expand":       true,
	}
	_, err := connections.XrplClient.Request(request)
	if err != nil {
		logger.Log.Error().Str("ledger_index", ledgerIndex).Msg(err.Error())
	}
}
