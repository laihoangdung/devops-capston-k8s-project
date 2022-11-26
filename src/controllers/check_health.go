package controllers

import "github.com/gin-gonic/gin"

func CheckHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "Ok"})
}