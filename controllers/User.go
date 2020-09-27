package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
)

// ListUsers list users
func ListUsers(c *gin.Context) {
	var users []models.User

	err := repositories.GetAllUsers(&users)

	if err != nil {
		c.JSON(404, "No user found")
	} else {
		c.JSON(200, users)
	}
}
