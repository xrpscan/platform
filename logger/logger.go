package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/xrpscan/platform/config"
)

var Log zerolog.Logger

func LoggerSetup() {
	level, err := zerolog.ParseLevel(config.EnvLogLevel())
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	zerolog.SetGlobalLevel(level)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Log = zerolog.New(output).With().Timestamp().Logger()
}
