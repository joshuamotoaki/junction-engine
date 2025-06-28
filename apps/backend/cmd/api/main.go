package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tigerappsorg/junction-engine/internal/api"
	"github.com/tigerappsorg/junction-engine/internal/shared/config"
)

// Main entry point for the Junction Engine API.
//
//	@title			Junction Engine API
//	@version		1.0
//	@description	This is your API description
//	@contact.name	TigerApps
//	@contact.url	https://tigerapps.org
//	@contact.email	it.admin@tigerapps.org
//	@license.name	BSD 3-Clause License
//	@license.url	https://opensource.org/licenses/BSD-3-Clause
//	@host			junction.tigerapps.org
//	@BasePath		/api/v1
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

	log.Info().Msgf("Log level set to %s", zerolog.GlobalLevel().String())

	api.Run(cfg)
}
