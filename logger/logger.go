package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xrpscan/platform/config"
)

var Log zerolog.Logger
var loggerOnce sync.Once

func New() {
	loggerOnce.Do(func() {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		level, err := zerolog.ParseLevel(config.EnvLogLevel())
		if err == nil {
			zerolog.SetGlobalLevel(level)
		}

		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		Log = zerolog.New(output).With().Timestamp().Logger()
	})
}
