package connections

import (
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/xrpl"
)

var XrplClient *xrpl.Client
var XrplFHClient *xrpl.Client

func NewXrplClient() {
	XrplClient = xrpl.NewClient(xrpl.ClientConfig{URL: config.EnvRippledURL()})
}

func NewXrplFHClient() {
	XrplFHClient = xrpl.NewClient(xrpl.ClientConfig{URL: config.EnvRippledFullHistoryURL()})
}
