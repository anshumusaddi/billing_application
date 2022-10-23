package routes

import (
	"github.com/anshumusaddi/billing_application/billing_event"
	"github.com/anshumusaddi/billing_application/models"
	"github.com/gin-gonic/gin"
)

func InitRoutes(events *[]models.MessageEvent) *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("billing_application/api/v1")

	messageEvent := baseRouter.Group("message/event")

	messageEvent.GET("/", func(c *gin.Context) { billing_event.GetBillingEvent(c, events) })
	messageEvent.POST("/", func(c *gin.Context) { billing_event.PostBillingEvent(c, events) })

	return router
}
