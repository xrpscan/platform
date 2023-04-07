package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/responses"
	"github.com/xrpscan/platform/xrpl"
)

func GetTransaction(c echo.Context) error {
	hash := c.Param("hash")
	producers.Produce(connections.KafkaWriter, xrpl.StreamMessage{Key: []byte(hash), Value: []byte(hash)})
	return c.JSON(http.StatusOK, responses.TransactionResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"tx": hash},
	})
}
