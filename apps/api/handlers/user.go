package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/database"
)

type UserHandler struct {
	db *database.Neo4jDB
}

func NewUserHandler(db *database.Neo4jDB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	// TODO: Implement list users in database module
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "List users not implemented yet",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// TODO: Implement update user
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Update user not implemented yet",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// TODO: Implement delete user
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Delete user not implemented yet",
	})
}
