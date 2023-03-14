package main

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/consumers"
	"github.com/xrpscan/platform/routes"
)

func main() {
	config.EnvLoad()

	connections.NewProducer()
	connections.NewConsumer()
	connections.NewEsClient()
	go consumers.RunTransactionConsumer()

	e := echo.New()
	routes.TransactionRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
