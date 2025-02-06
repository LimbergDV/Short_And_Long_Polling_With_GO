package controllers

import (
	"api_short_long_polling/src/cars/application"
	"api_short_long_polling/src/cars/domain"
	"api_short_long_polling/src/cars/infrastructure"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCarByIdController struct {
	app *application.UpdateCar
}

func NewUpdateCarByIdController() *UpdateCarByIdController {
	mysql := infrastructure.GetMySQL()
	app := application.NewUpdateCar(mysql)
	return &UpdateCarByIdController{app: app}
}

func (ctrl *UpdateCarByIdController) Run(c *gin.Context) {
	id := c.Param("id")
	var car domain.Car

	idCar, _ := strconv.ParseUint(id, 10, 64)

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := ctrl.app.Run(int(idCar), car)

	if rowsAffected == 0 {
		fmt.Println("No se pudo actualizar el carro")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el carro"})
		return
	}

	// Enviar una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Car updated successfully",
	})
}
