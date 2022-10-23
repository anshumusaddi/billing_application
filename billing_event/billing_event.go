package billing_event

import (
	"github.com/anshumusaddi/billing_application/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBillingEvent(c *gin.Context, events *[]models.MessageEvent) {
	c.JSON(http.StatusOK, events)
}

func PostBillingEvent(c *gin.Context, events *[]models.MessageEvent) {
	messageEvent := &models.MessageEvent{}
	err := c.ShouldBindJSON(messageEvent)
	if err != nil {
		println("Error parsing request body, err: %s", err.Error())
		return
	}
	*events = append(*events, *messageEvent)
	c.JSON(http.StatusOK, events)
}
