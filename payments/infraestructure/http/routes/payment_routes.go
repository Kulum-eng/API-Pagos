package routes

import (
	"github.com/gin-gonic/gin"

	"ModaVane/payments/infraestructure/http/controllers"

)

func SetupPaymentRoutes(router *gin.Engine, controller *controllers.PaymentController) {
	paymentRoutes := router.Group("/payments")
	{
		paymentRoutes.POST("/", controller.Create)
		paymentRoutes.GET("/", controller.GetAll)
		paymentRoutes.GET("/:id", controller.GetByID)
		paymentRoutes.PUT("/:id", controller.Update)
		paymentRoutes.DELETE("/:id", controller.Delete)
	}
}
