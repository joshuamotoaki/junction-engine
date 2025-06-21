package application

import (
	"context"
	"log"

	"github.com/tigerappsorg/junction-engine/auth"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
)

func Run(cfg *config.Config) {
	// Neo4J
	db, err := database.NewNeo4j(cfg.NEO4J_URI, cfg.NEO4J_USERNAME, cfg.NEO4J_PASSWORD)
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}
	defer db.Close(context.Background())

	// CAS Authentication
	casService := auth.NewCASService(cfg)

	// Router
	router := NewRouter(cfg, db, casService)
	router.SetupRoutes()
	router.Run()
}
