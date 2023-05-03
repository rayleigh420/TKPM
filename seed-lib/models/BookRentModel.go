package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRentModel struct {
	Id				primitive.ObjectID	`bson:"_id"`
	User_id			string				`json:"user_id"`
	Book_id			string				`json:"book_id"`
	Book_detail_id	string				`json:"book_detail_id"`
	Reserve_date 	time.Time			`json:"reserve_date"`
	Book_rent_id	string				`json:"book_rent_id"`
}