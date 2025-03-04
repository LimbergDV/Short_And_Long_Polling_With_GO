package controllers

import (
	"api_short_long_polling/src/customers/application"
	"api_short_long_polling/src/customers/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)


func ShortPollingCustomers(c *gin.Context) {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetAllCustomers(mysql)
	c.Writer.Flush()
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			customersData := useCase.Run()
			// Creamos el mensaje en formato JSON
			data, err := json.Marshal(gin.H{
				"message":   "Datos actuales de clientes",
				"customers": customersData,
			})
			if err != nil {
				fmt.Println("Error formateando JSON:", err)
				continue
			}
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()

		case <-c.Writer.CloseNotify():
			return
		}
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

