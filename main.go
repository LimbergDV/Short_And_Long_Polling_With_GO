package main

import (
	"api_short_long_polling/src/cars/infrastructure"
	routesCars "api_short_long_polling/src/cars/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	infrastructure.GoMySQL()

	r := gin.Default()

	routesCars.Routes(r)
	r.Run(":8081")
}