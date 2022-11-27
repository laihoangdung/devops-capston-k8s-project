package controllers

import (
	"venv/src/models"

	"github.com/gin-gonic/gin"
)

// Get list of users
func GetAllUsers(ctx *gin.Context) {
	users := []models.User{
		{
			Id: 1,
			Name: "Tom",
			Age: 20,
		},
		{
			Id: 2,
			Name: "Jerry",
			Age: 21,
		},
		{
			Id: 3,
			Name: "Ronaldo",
			Age: 22,
		},
		{
			Id: 4,
			Name: "Messi",
			Age: 23,
		},
		{
			Id: 5,
			Name: "Kaka",
			Age: 24,
		},
		{
			Id: 6,
			Name: "Mbappe",
			Age: 25,
		},
	}

	ctx.JSON(200, gin.H{
		"users": users,
	})
}