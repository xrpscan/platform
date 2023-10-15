package connections

import (
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/xrpl-go"
)

var XrplClient *xrpl.Client

func NewXrplClient() {
	NewXrplClientWithURL(config.EnvXrplWebsocketURL())
}

func NewXrplClientWithURL(URL string) {
	XrplClient = xrpl.NewClient(xrpl.ClientConfig{URL: URL})
	err := XrplClient.Ping([]byte(URL))
	if err != nil {
		panic(err)
	}
}
