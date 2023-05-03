package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookModel struct {
	Id                  primitive.ObjectID `bson:"_id"`
	Name                string             `json:"name"`
	Publisher           string             `json:"publisher"`
	Yearpublished       int64              `json:"yearpublished"`
	Author              string             `json:"author"`
	Book_id             string             `json:"book_id"`
	Book_img            string             `json:"book_image"`
	Amount              int64              `json:"amount"`
	Type_id             string             `json:"type_id"`
	Page                int64              `json:"page"`
	Publishing_location string             `json:"publishing_location"`
	Borrowed_quantity   int64              `json:"borrowed_quantity"`
	Details             string             `json:"details"`
	Description         string             `json:"description"`
	License             string             `json:"license"`
	Created_at          time.Time          `json:"created_at"`
	Updated_at          time.Time          `json:"updated_at"`
}
