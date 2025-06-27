package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/tigerappsorg/junction-engine/internal/shared/models"
)

func (db *neo4jDB) CreateUser(ctx context.Context, user *models.User) error {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
            MERGE (u:User {netid: $netid})
            SET u.name = $name,
                u.email = $email,
                u.class_year = $class_year,
                u.updated_at = datetime(),
                u.created_at = CASE 
                    WHEN u.created_at IS NULL THEN datetime() 
                    ELSE u.created_at 
                END
            RETURN u
        `
		_, err := tx.Run(ctx, query, map[string]any{
			"netid":      user.NetID,
			"name":       user.Name,
			"email":      user.Email,
			"class_year": user.ClassYear,
		})
		return nil, err
	})

	return err
}

func (db *neo4jDB) GetUserByNetID(ctx context.Context, netid string) (*models.User, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
            MATCH (u:User {netid: $netid})
            RETURN u.netid as netid, 
                   u.name as name, 
                   u.email as email, 
                   u.class_year as class_year,
                   u.created_at as created_at
        `
		result, err := tx.Run(ctx, query, map[string]any{
			"netid": netid,
		})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			record := result.Record()
			user := &models.User{
				NetID:     record.AsMap()["netid"].(string),
				Name:      record.AsMap()["name"].(string),
				Email:     record.AsMap()["email"].(string),
				ClassYear: record.AsMap()["class_year"].(string),
			}
			return user, nil
		}

		return nil, nil // No user found
	})

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result.(*models.User), nil
}
