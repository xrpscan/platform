package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/responses"
	"github.com/xrpscan/platform/xrpl"
)

func GetAccountInfo(c echo.Context) error {
	address := c.Param("address")
	producers.Produce(connections.KafkaWriter, []byte(address))
	req := xrpl.BaseRequest{
		"command": "account_info",
		"account": "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg",
	}
	res, _ := connections.XrplClient.Request(req)
	return c.JSON(http.StatusOK, responses.TransactionResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"account_info": address},
		Result:  &res,
	})
}
