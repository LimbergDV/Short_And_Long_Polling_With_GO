package controllers

import (
	"api_short_long_polling/src/cars/application"
	"api_short_long_polling/src/cars/infrastructure"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// ShortPollingAvailableCars responde inmediatamente con la lista de carros disponibles.
func ShortPollingAvailableCars(c *gin.Context) {
	mysql := infrastructure.GetMySQL() // Obtenemos la instancia de MySQL.
	useCase := application.NewGetAvailableCars(mysql)
	availableCars := useCase.Run()
	c.JSON(http.StatusOK, availableCars)
}

// LongPollingAvailableCars mantiene la conexión abierta hasta detectar un cambio.
func LongPollingAvailableCars(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAvailableCars(mysql)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	initialCars := useCase.Run()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
			return
		case <-ticker.C:
			updatedCars := useCase.Run()
			if !reflect.DeepEqual(initialCars, updatedCars) {
				c.JSON(http.StatusOK, updatedCars)
				return
			}
		}
	}
}
