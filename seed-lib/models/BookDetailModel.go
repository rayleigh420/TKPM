package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BookDetailModel struct {
	Id             primitive.ObjectID `bson:"_id"`
	Book_id        string             `json:"book_id"`
	Book_detail_id string             `json:"book_detail_id"`
	Location       string             `json:"location"`
	Status         string             `json:"status"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}