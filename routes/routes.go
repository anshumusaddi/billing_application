package routes

import (
	"github.com/anshumusaddi/billing_application"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	baseRouter := router.Group("billing_application/api/v1")

	billingEvent := baseRouter.Group("billing/event")

	billingEvent.GET("/", func(c *gin.Context) { getBillingEvent(c) })

}
