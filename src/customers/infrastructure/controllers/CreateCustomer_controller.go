package controllers

import (
	"api_short_long_polling/src/customers/application"
	"api_short_long_polling/src/customers/domain"
	"api_short_long_polling/src/customers/infrastructure"
	"api_short_long_polling/src/customers/infrastructure/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCustomerController struct {
	app *application.CreateCustomer
}

func NewCreateCustomerController() *CreateCustomerController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreateCustomer(mysql)
	return &CreateCustomerController{app: app}
}

func (cc_c *CreateCustomerController) Run (c *gin.Context){
	var customers domain.Customer

	if err := c.ShouldBindJSON(&customers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Datos inválidos" + err.Error()})
		return
	}

	if err := validators.CheckCustomer(customers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos invalidos" + err.Error()})
	}
	
	rowsAffected, err := cc_c.app.Run(customers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if rowsAffected == 0{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Cliente creado"})
		c.JSON(http.StatusOK, customers)
	}
}