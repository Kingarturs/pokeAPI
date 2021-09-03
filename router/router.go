package router

import (
	"pokeAPI/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var BaseRoute = gin.Default().BasePath()

func Routes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/hello", routes.HelloWorld())
	router.GET("/", routes.Pokemon())

	return router
}
