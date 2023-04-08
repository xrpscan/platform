package connections

import (
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/xrpl"
)

var XrplClient *xrpl.Client

func NewXrplClient() {
	XrplClient = xrpl.NewClient(xrpl.ClientConfig{URL: config.EnvRippledURL()})
	err := XrplClient.Ping([]byte(config.EnvRippledURL()))
	if err != nil {
		panic(err)
	}
}
