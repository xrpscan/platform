package responses

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/xrpl"
)

type TransactionResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *echo.Map          `json:"data"`
	Result  *xrpl.BaseResponse `json:"result"`
}
