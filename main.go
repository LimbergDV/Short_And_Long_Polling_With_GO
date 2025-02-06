package main

import (
	cars "api_short_long_polling/src/cars/infrastructure"
	routesCars "api_short_long_polling/src/cars/infrastructure/routes"
	customers "api_short_long_polling/src/customers/infrastructure"
	routesCustomers"api_short_long_polling/src/customers/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	cars.GoMySQL()
	customers.GoMySQL()

	r := gin.Default()

	routesCars.Routes(r)
	routesCustomers.Routes(r)
	
	r.Run(":8081")
}