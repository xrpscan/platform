package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/controllers"
)

func TransactionRoute(e *echo.Echo) {
	e.GET("/tx/:hash", controllers.GetTransaction)
}
