package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/responses"
)

func GetTransaction(c echo.Context) error {
	hash := c.Param("hash")
	producers.Produce(connections.KafkaProducer, hash)
	return c.JSON(http.StatusOK, responses.TransactionResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"tx": hash},
	})
}
