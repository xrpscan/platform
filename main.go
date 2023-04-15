package main

import (
	"github.com/labstack/echo/v4"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/consumers"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/routes"
	"github.com/xrpscan/platform/signals"
)

func main() {
	config.EnvLoad()

	connections.NewWriter()
	connections.NewReaders()
	connections.NewEsClient()
	connections.NewXrplClient()

	producers.SubscribeStreams()
	consumers.RunConsumers()

	e := echo.New()
	e.HideBanner = true
	routes.Add(e)

	signals.HandleAll()
	e.Logger.Fatal(e.Start(":3000"))
}
