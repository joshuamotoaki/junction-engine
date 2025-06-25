package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/database"
)

type userHandler struct {
	db database.Neo4jDB
}

type UserHandler interface {
	ListUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler(db database.Neo4jDB) UserHandler {
	return &userHandler{
		db: db,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *userHandler) ListUsers(c *gin.Context) {
	// TODO: Implement list users in database module
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "List users not implemented yet",
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	// TODO: Implement update user
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Update user not implemented yet",
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	// TODO: Implement delete user
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Delete user not implemented yet",
	})
}
