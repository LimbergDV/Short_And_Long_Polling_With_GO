package controllers

import (
	"api_short_long_polling/src/cars/application"
	"api_short_long_polling/src/cars/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCarByIdController struct {
	app *application.DeleteCar
}

func NewDeleteCarByIdController() *DeleteCarByIdController {
	mysql := infrastructure.GetMySQL()
	app := application.NewDeleteCar(mysql)
	return &DeleteCarByIdController{app: app}
}

func (ctrl *DeleteCarByIdController) Run(c *gin.Context) {
	id := c.Param("id")
	carId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del carro inválido"})
		return
	}

	rowsAffected, _ := ctrl.app.Run(carId)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el carro"})
		return
	}

	// Devolviendo el mensaje y el ID eliminado
	c.JSON(http.StatusOK, gin.H{"message": "Carro eliminado exitosamente"})
}
