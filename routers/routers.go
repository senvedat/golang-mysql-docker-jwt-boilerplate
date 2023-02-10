package routers

import (
	"go-rest-example/controllers"
	"go-rest-example/routers/middleware"

	"github.com/gin-gonic/gin"
)

func MainRoutes(route *gin.Engine) {
	spaceship := route.Group("/spaceship")
	spaceship.GET("/list", middleware.JWTMiddleware(), controllers.GetSpaceships)
	spaceship.POST("/create", middleware.JWTMiddleware(), controllers.CreateSpaceship)
	spaceship.GET("/:id", middleware.JWTMiddleware(), controllers.GetSpaceship)
	spaceship.PUT("/:id", middleware.JWTMiddleware(), controllers.UpdateSpaceship)
	spaceship.DELETE("/:id", middleware.JWTMiddleware(), controllers.DeleteSpaceship)

	login := route.Group("/login")
	login.POST("/", controllers.Login)
}
