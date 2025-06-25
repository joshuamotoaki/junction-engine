package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tigerappsorg/junction-engine/application"
	"github.com/tigerappsorg/junction-engine/config"
)

func main() {
	cfg := config.Load()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if cfg.Env == "dev" {
		if cfg.Debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		// Pretty logger for development
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)

		// JSON logger for production
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	application.Run(cfg)
}
