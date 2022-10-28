package helper

import (
	"github.com/anshumusaddi/billing_application/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"strconv"
	"time"
)

func TimeAddr(i time.Time) *time.Time {
	timeVar := i
	return &timeVar
}

func GetMapFromQueryParams(queryParams url.Values) map[string]string {
	queryMap := make(map[string]string)
	for k := range queryParams {
		queryMap[k] = queryParams.Get(k)
	}
	return queryMap
}

func GetBsonFromQuery(queryParams map[string]string, queryToBsonMap map[string]string) (bson.M, *APIError) {
	bsonQuery := bson.M{}
	for k, v := range queryParams {
		if k == "id" {
			bsonQuery[queryToBsonMap[k]], _ = primitive.ObjectIDFromHex(v)
			continue
		}
		if k == "size" {
			bsonQuery[queryToBsonMap[k]], _ = strconv.ParseInt(v, 10, 64)
			continue
		}
		key, ok := queryToBsonMap[k]
		if !ok {
			return nil, ErrInvalidQueryParams
		}
		bsonQuery[key] = v
	}
	logger.Debug("the find query getting executed is : ", bsonQuery)
	return bsonQuery, nil
}

func GetDateEqualityExpression(part string, field string, value int) bson.D {
	extractedSection := bson.D{{"$" + part, "$" + field}}
	bsonQuery := bson.D{{"$eq", []interface{}{
		extractedSection,
		value,
	}}}
	return bsonQuery
}
