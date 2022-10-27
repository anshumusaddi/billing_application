package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

func ObjectIDAddr(i primitive.ObjectID) *primitive.ObjectID {
	ObjectIDVar := i
	return &ObjectIDVar
}

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
		key, ok := queryToBsonMap[k]
		if !ok {
			return nil, ErrInvalidQueryParams
		}
		bsonQuery[key] = v
	}
	return bsonQuery, nil
}
