package main

import (
	"github.com/gin-gonic/gin"
	"venv/src/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/check-health", controllers.CheckHealth)
	users := r.Group("/users")
	{
		users.GET("", controllers.GetAllUsers)
	}
	
	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}