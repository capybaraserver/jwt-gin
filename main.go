package main

import (
	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()
	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddlewares())
	api.GET("/user", controllers.GetUser)

	r.Run(":8080")

}
