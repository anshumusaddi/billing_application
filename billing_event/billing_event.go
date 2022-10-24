package billing_event

import (
	"github.com/anshumusaddi/billing_application/helper"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/anshumusaddi/billing_application/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBillingEvent(ctx *gin.Context, events *[]models.MessageEvent) {
	helper.WriteSuccessResponse(ctx, http.StatusOK, events)
}

func PostBillingEvent(ctx *gin.Context, events *[]models.MessageEvent) {
	messageEvent := &models.MessageEvent{}
	err := ctx.ShouldBindJSON(messageEvent)
	if err != nil {
		logger.Error("error parsing request body, err: %s", err.Error())
		helper.WriteErrorResponse(ctx, helper.ErrInvalidRequestPayloadParams)
		return
	}
	*events = append(*events, *messageEvent)
	helper.WriteSuccessResponse(ctx, http.StatusOK, events)
}
