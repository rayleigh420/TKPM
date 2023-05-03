package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryModel struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	User_id        string             `json:"user_id"`
	Book_id        string             `json:"book_id"`
	Book_detail_id string             `json:"book_detail_id"`
	Status         string             `json:"status"`
	Date_borrowed  time.Time          `json:"date_borrowed"`
	Date_return    time.Time          `json:"date_return"`
	History_id     string             `json:"history_id"`
}
