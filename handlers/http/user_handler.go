package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler is a handler for creating a user
func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

// ListUserHandler is a handler for listing all users
func ListUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// GetUserHandler is a handler for getting a user
func GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// UpdateUserHandler is a handler for updating a user
func UpdateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// DeleteUserHandler is a handler for deleting a user
func DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
