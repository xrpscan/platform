package main

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/routes"
)

func main() {
	e := echo.New()
	connections.NewProducer()
	connections.NewConsumer()

	routes.TransactionRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
