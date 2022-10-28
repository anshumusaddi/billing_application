package models

import "time"

type SummaryGroupID struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	Month      int64  `json:"month" bson:"month"`
	Year       int64  `json:"year" bson:"year"`
}

type SummaryInfo struct {
	ID       SummaryGroupID `json:"id" bson:"_id"`
	Total    int64          `json:"total" bson:"total"`
	LastDate time.Time      `json:"last_date" bson:"last_date"`
}
