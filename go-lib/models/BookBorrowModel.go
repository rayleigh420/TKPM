package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookBorrowModel struct {
	Id  			primitive.ObjectID	`bson:"_id"`
	User_id			string				`json:"user_id"`
	Book_id			string				`json:"book_id"`
	Book_detail_id	string				`json:"book_detail_id"`
	Date_borrowed	time.Time			`json:"date_borrowed"`
	Date_end		time.Time			`json:"date_end"`
	Book_hire_id	string				`json:"book_hire_id"`
}


