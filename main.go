package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	p_application "ModaVane/payments/application"
	core "ModaVane/payments/core"
	p_adapters "ModaVane/payments/infraestructure/adapters"
	p_controllers "ModaVane/payments/infraestructure/http/controllers"
	p_routes "ModaVane/payments/infraestructure/http/routes"

)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	myGin := gin.New()
	myGin.RedirectTrailingSlash = false

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println(err)
		return
	}

	paymentRepository := p_adapters.NewMySQLPaymentRepository(db)
	senderNotification := p_adapters.NewHTTPSenderNotification("localhost", 3000)
	

	createPaymentUseCase := p_application.NewCreatePaymentUseCase(paymentRepository, senderNotification)
	getPaymentUseCase := p_application.NewGetPaymentUseCase(paymentRepository)
	updatePaymentUseCase := p_application.NewUpdatePaymentUseCase(paymentRepository)
	deletePaymentUseCase := p_application.NewDeletePaymentUseCase(paymentRepository)

	createPaymentController := p_controllers.NewPaymentController(createPaymentUseCase, getPaymentUseCase, updatePaymentUseCase, deletePaymentUseCase)
	p_routes.SetupPaymentRoutes(myGin, createPaymentController)

	myGin.Run(":8081")
}
