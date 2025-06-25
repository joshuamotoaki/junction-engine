package database

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/config"
	"github.com/tigerappsorg/junction-engine/models"
)

type neo4jDB struct {
	Driver neo4j.DriverWithContext
}

type Neo4jDB interface {
	Close(ctx context.Context) error
	HealthCheck(ctx context.Context) error

	CreateUser(ctx context.Context, user *models.User) error
}

func NewNeo4j(uri, username, password string) (Neo4jDB, error) {
	driver, err := neo4j.NewDriverWithContext(
		uri,
		neo4j.BasicAuth(username, password, ""),
		func(config *config.Config) {
			config.MaxConnectionLifetime = 0
			config.MaxConnectionPoolSize = 50
			config.ConnectionAcquisitionTimeout = 0
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	ctx := context.Background()
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		err := driver.Close(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to close driver after connectivity check failure: %w", err)
		}
		return nil, fmt.Errorf("failed to verify connectivity: %w", err)
	}

	return &neo4jDB{Driver: driver}, nil
}

func (db *neo4jDB) Close(ctx context.Context) error {
	if db.Driver != nil {
		return db.Driver.Close(ctx)
	}
	return nil
}

func (db *neo4jDB) HealthCheck(ctx context.Context) error {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.Run(ctx, "RETURN 1", nil)
	return err
}
