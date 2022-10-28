package billing_event

import (
	"errors"
	"github.com/anshumusaddi/billing_application/datastore"
	"github.com/anshumusaddi/billing_application/helper"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/anshumusaddi/billing_application/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func GetBillingEvent(ctx *gin.Context, store *datastore.BillingApplicationDBStore) {
	var messagingEvents []models.MessageEvent
	queryParams := helper.GetMapFromQueryParams(ctx.Request.URL.Query())
	query, apiErr := helper.GetBsonFromQuery(queryParams, billingEventQueryMap)
	if apiErr != nil {
		logger.Error("error parsing query params, err: %s", apiErr.Error())
		helper.WriteErrorResponse(ctx, apiErr)
		return
	}
	err := store.Find(datastore.MessagingEventCollection, &messagingEvents, query)
	if err != nil {
		logger.Error("error loading messaging events from database, err: %s", err.Error())
		helper.WriteErrorResponse(ctx, helper.ApiErrorWithCustomMessage(helper.ErrDBOperation, err.Error()))
		return
	}
	helper.WriteSuccessResponse(ctx, http.StatusOK, messagingEvents)
}

func PostBillingEvent(ctx *gin.Context, store *datastore.BillingApplicationDBStore) {
	messageEvent := &models.MessageEvent{}
	err := ctx.ShouldBindJSON(&messageEvent)
	if err != nil {
		logger.Error("error parsing request body, err: %s", err.Error())
		helper.WriteErrorResponse(ctx, helper.ErrInvalidRequestPayloadParams)
		return
	}
	if messageEvent.CustomerID == "" || messageEvent.Size == 0 {
		err = errors.New("mandatory fields not present in input")
		logger.Error("error parsing request body, err: %s", err.Error())
		helper.WriteErrorResponse(ctx, helper.ErrInvalidRequestPayloadParams)
		return
	}
	messageEvent.ID = primitive.NewObjectID()
	timeStamp := time.Now()
	messageEvent.CreatedAt = helper.TimeAddr(timeStamp)
	messageEvent.UpdatedAt = helper.TimeAddr(timeStamp)
	messageEvent.EventTime = timeStamp
	err = store.CreateOne(datastore.MessagingEventCollection, messageEvent)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			logger.Error("unique constrain violates for messaging_event collection")
			helper.WriteErrorResponse(ctx, helper.ApiErrorWithCustomMessage(helper.ErrDuplicateKey, err.Error()))
		} else {
			logger.Error("error persisting to database, err: %s", err.Error())
			helper.WriteErrorResponse(ctx, helper.ApiErrorWithCustomMessage(helper.ErrDBInsert, err.Error()))
		}
		return
	}
	helper.WriteSuccessResponse(ctx, http.StatusOK, messageEvent)
	return
}
