package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/controllers"
)

func Add(e *echo.Echo) {
	e.GET("/tx/:hash", controllers.GetTransaction)
	e.GET("/account/:address", controllers.GetAccountInfo)
}
