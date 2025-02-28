package controllers

import (
	"api_short_long_polling/src/customers/application"
	"api_short_long_polling/src/customers/infrastructure"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// ShortPollingCustomers responde de forma inmediata con el estado actual de los customers.
func ShortPollingCustomers(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAllCustomers(mysql)
	
	ticket := time.NewTicker(15 * time.Second)


	for range ticket.C {
		customersData := useCase.Run()
		c.JSON(http.StatusOK, gin.H{
			"message": "Datos actuales de clientes",
			"cars":    customersData,
		})
	}
}

// LongPollingCustomers mantiene la conexión abierta hasta detectar cualquier cambio en los atributos de los customers.
func LongPollingCustomers(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAllCustomers(mysql)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	initialCustomers := useCase.Execute()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
			return
		case <-ticker.C:
			updatedCustomers := useCase.Execute()
			// Se utiliza reflect.DeepEqual para comparar todos los atributos de cada customer.
			if !reflect.DeepEqual(initialCustomers, updatedCustomers) {
				c.JSON(http.StatusOK, gin.H{
					"message":   "Se detectaron cambios en los customers",
					"customers": updatedCustomers,
				})
				return
			}
		}
	}
}

