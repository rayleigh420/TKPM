package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookTypeModel struct {
	Id       primitive.ObjectID `bson:"_id"`
	Type_id   string 			`json:"type_id"`
	Type_name string 			`json:"type_name"`
}