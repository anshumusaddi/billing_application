package routes

import (
	"github.com/anshumusaddi/billing_application/billing_event"
	"github.com/anshumusaddi/billing_application/billing_summary"
	"github.com/anshumusaddi/billing_application/datastore"
	"github.com/gin-gonic/gin"
)

func InitRoutes(store *datastore.BillingApplicationDBStore, kafkaStore *datastore.BillingApplicationKafkaStore) *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("billing_application/api/v1")

	messageEvent := baseRouter.Group("message/event")

	messageEvent.GET("/", func(c *gin.Context) { billing_event.GetBillingEvent(c, store) })
	messageEvent.POST("/", func(c *gin.Context) { billing_event.PostBillingEvent(c, kafkaStore) })

	messageSummary := baseRouter.Group("message/summary")

	messageSummary.GET("/", func(c *gin.Context) { billing_summary.GetBillingSummary(c, store) })

	return router
}
