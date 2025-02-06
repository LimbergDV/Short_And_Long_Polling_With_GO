package controllers

import (
	"api_short_long_polling/src/customers/application"
	"api_short_long_polling/src/customers/domain"
	"api_short_long_polling/src/customers/infrastructure"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCustomerByIdController struct{
	app *application.UpdateCustomer
}

func NewUpdateCustomerByIdController() *UpdateCustomerByIdController{
	mysql := infrastructure.GetMySQL()
	app := application.NewUpdateCustomer(mysql)
	return &UpdateCustomerByIdController{app: app}
}

func (ctrl *UpdateCustomerByIdController) Run(c *gin.Context){
	id := c.Param("id")
	var car domain.Customer

	id_cutomer, _ := strconv.ParseUint(id, 10, 64)

	if err :=  c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	RowsAffected, _ := ctrl.app.Run(int(id_cutomer), car)

	if RowsAffected == 0{
		fmt.Print("hola")
	}

	// Send a successful response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Customer updated successfully",
	})

}