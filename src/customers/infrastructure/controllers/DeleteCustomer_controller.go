package controllers

import (
	"api_short_long_polling/src/customers/application"
	"api_short_long_polling/src/customers/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCustomerByIdController struct{
	app *application.DeleteCustomer
}

func NewDeleteCustomerByIdController() *DeleteCustomerByIdController {
	mysql := infrastructure.GetMySQL()
	app := application.NewDeleteCustomer(mysql)
	return &DeleteCustomerByIdController{app: app}
}

func (ctrl *DeleteCustomerByIdController) Run(c *gin.Context)  {
	id := c.Param("id")
	customerId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id de cliente invalido"})
		return
	}
	rowsAffected, _ := ctrl.app.Run(customerId)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}

	// Devolviendo el mensaje y el id eliminado
	c.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado exitosamente"})
}