package models

type MessageEvent struct {
	ID         *string `json:"id" bson:"_id"`
	CreatedAt  *int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  *int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt  *int64  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	CustomerID *string `json:"customer_id" bson:"customer_id"`
	Size       *int64  `json:"size" bson:"size"`
}
