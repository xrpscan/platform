package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/responses"
)

func GetAccountInfo(c echo.Context) error {
	address := c.Param("address")
	producers.Produce(connections.KafkaWriter, []byte(address))
	req := fmt.Sprintf("{\"id\": 1,\"command\": \"account_info\", \"account\": \"%s\"}", address)
	fmt.Println("req:", req)
	err := connections.XrplClient.Request([]byte(req))
	if err != nil {
		fmt.Println("Error sending account_info")
	}
	return c.JSON(http.StatusOK, responses.TransactionResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"account_info": address},
	})
}
