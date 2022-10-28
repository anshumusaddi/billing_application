package billing_summary

import (
	"github.com/anshumusaddi/billing_application/datastore"
	"github.com/anshumusaddi/billing_application/helper"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

func GetBillingSummary(ctx *gin.Context, store *datastore.BillingApplicationDBStore) {
	matchStage := bson.D{}
	groupStage := bson.D{}
	matchElements := bson.D{}
	queryParams := helper.GetMapFromQueryParams(ctx.Request.URL.Query())
	customerId, ok := queryParams["customer_id"]
	if ok {
		matchElements = append(matchElements, bson.E{Key: "customer_id", Value: customerId})
	}
	var andElements []interface{}
	month, ok := queryParams["month"]
	if ok {
		month, _ := strconv.Atoi(month)
		andElements = append(andElements, helper.GetDateEqualityExpression("month", "date", month))
	}
	year, ok := queryParams["year"]
	if ok {
		year, _ := strconv.Atoi(year)
		andElements = append(andElements, helper.GetDateEqualityExpression("year", "date", year))
	}
	matchElements = append(matchElements, bson.E{Key: "$expr", Value: bson.E{Key: "$and", Value: andElements}})
	store.RemoveDeletedDocuments(matchElements)
	matchStage = append(matchStage, bson.E{Key: "$match", Value: matchElements})
	logger.Debug("match stage constructed : ", matchStage)
	groupStage = append(groupStage, bson.E{Key: "$group", Value: bson.D{
		{"_id", bson.D{
			{"customer_id", "$customer_id"},
			{"month", bson.D{
				{"$month", "$time"}}},
			{"year", bson.D{
				{"$year", "$time"}}},
		}},
		{"total", bson.D{{"$sum", "$size"}}},
	}})
	logger.Debug("group stage constructed : ", groupStage)
	pipeline := mongo.Pipeline{matchStage, groupStage}
	var aggregateInfo []bson.M
	err := store.Aggregate(datastore.MessagingEventCollection, &aggregateInfo, pipeline)
	if err != nil {
		logger.Error("error executing aggregation pipeline in database, err: %s", err.Error())
		helper.WriteErrorResponse(ctx, helper.ApiErrorWithCustomMessage(helper.ErrDBOperation, err.Error()))
		return
	}
	helper.WriteSuccessResponse(ctx, http.StatusOK, aggregateInfo)
	return
}
