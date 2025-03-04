package controllers

import (
	"api_short_long_polling/src/cars/application"
	"api_short_long_polling/src/cars/domain"
	"api_short_long_polling/src/cars/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func ShortPollingAvailableCars(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAvailableCars(mysql)
	c.Writer.Flush()
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			availableCars := useCase.Run()
			// Creamos el mensaje en formato JSON
			data, err := json.Marshal(gin.H{
				"message": "Datos actuales de disponibilidad",
				"cars":    availableCars,
			})
			if err != nil {
				fmt.Println("Error al formatear JSON:", err)
				continue
			}

			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
			
		case <-c.Writer.CloseNotify():
			return
		}
	}
}

// Función para comparar solo el campo Available de dos slices de Car
func hasAvailabilityChanged(initial, updated []domain.Car) bool {
	if len(initial) != len(updated) {
		return true
	}
	for i := range initial {
		if initial[i].Available != updated[i].Available {
			return true
		}
	}
	return false
}

// LongPollingAvailableCars mantiene la conexión abierta hasta detectar cambios en "Available"
func LongPollingAvailableCars(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAvailableCars(mysql)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	initialCars := useCase.Run()

	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
			return
		case <-ticker.C:
			updatedCars := useCase.Run()
			// Usamos la función custom para verificar solo el campo Available
			if hasAvailabilityChanged(initialCars, updatedCars) {
				c.JSON(http.StatusOK, gin.H{"message": "La disponibilidad de los carros ha cambiado"})
				c.JSON(http.StatusOK, gin.H{"cars": updatedCars})
				return
			}
		}
	}
}
//hacer perticiones varias veces con un for 
