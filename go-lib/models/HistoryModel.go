package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryModel struct  {
	Id 				primitive.ObjectID `bson:"_id,omitempty"`
	Status			string				`json:"status"`
	Date_return		time.Time			`json:"date_return"`
	History_id		string				`json:"history_id"`
	Book_hire_id	string				`json:"book_hire_id"`
}