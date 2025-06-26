package application

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tigerappsorg/junction-engine/auth"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
)

func Run(cfg *config.Config) {
	// Neo4J
	db, err := database.NewNeo4j(cfg.Neo4jURI, cfg.Neo4jUsername, cfg.Neo4jPassword)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to Neo4j:", err)
	}

	defer db.Close(context.Background())

	// CAS Authentication
	casService := auth.NewCASService(cfg)

	// Router
	router := NewRouter(cfg, db, casService)
	router.SetupRoutes()
	router.Run()
}
