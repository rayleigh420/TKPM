package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookTypeModel struct {
	Id       primitive.ObjectID `bson:"_id"`
	TypeId   string `json:"type_id"`
	TypeName string `json:"type_name"`
}