package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tigerappsorg/junction-engine/internal/database/neo4j"
	"github.com/tigerappsorg/junction-engine/internal/shared/auth"
	"github.com/tigerappsorg/junction-engine/internal/shared/config"
)

func Run(cfg *config.Config) {
	// Neo4J
	db, err := neo4j.NewNeo4j(cfg.Neo4jURI, cfg.Neo4jUsername, cfg.Neo4jPassword)
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
