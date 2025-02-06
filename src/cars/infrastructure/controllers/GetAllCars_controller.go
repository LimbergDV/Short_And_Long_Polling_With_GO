package controllers

import (
	"api_short_long_polling/src/cars/application"
	"api_short_long_polling/src/cars/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllCarsController struct {
	app *application.GetAllCars
}

func NewGetAllCarsController() *GetAllCarsController {
	mysql := infrastructure.GetMySQL()
	app := application.NewGetAllCars(mysql)
	return &GetAllCarsController{app: app}
}

func (ctrl *GetAllCarsController) Run(c *gin.Context) {
	res := ctrl.app.Run()

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "No se encontró ningún carro"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Carros": res})
	}
}