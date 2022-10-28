package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageEvent struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt  *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt  *time.Time         `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time         `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	EventTime  time.Time          `json:"time" bson:"time"`
	CustomerID string             `json:"customer_id" bson:"customer_id"`
	Size       int64              `json:"size" bson:"size"`
}
