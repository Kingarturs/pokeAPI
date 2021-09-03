package main

import (
	"fmt"
	"log"

	"pokeAPI/config"
	"pokeAPI/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(
			"\n",
			"\x1b[31mError loading .env file\x1b[0m\n",
			"\x1b[36mCopy the .default.env into a .env file\x1b[0m",
		)
	}

	if config.GetGinMode() == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := router.Routes()
	router.Run(fmt.Sprintf(":%s", config.GetPort()))
}
