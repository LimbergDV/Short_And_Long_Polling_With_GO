package routes

import (
	"api_short_long_polling/src/cars/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine) {
	
	carsRoutes := r.Group("/cars") 
	{
		carsRoutes.POST("/", controllers.NewCreateCarController().Run)
		carsRoutes.GET("/", controllers.NewGetAllCarsController().Run)
		carsRoutes.PUT("/:id", controllers.NewUpdateCarByIdController().Run)
		carsRoutes.DELETE("/:id", controllers.NewDeleteCarByIdController().Run)
		carsRoutes.GET("/available", controllers.ShortPollingAvailableCars)
		carsRoutes.GET("/available-wait", controllers.LongPollingAvailableCars)

	}
}