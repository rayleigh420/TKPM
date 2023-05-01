package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookTypeModel struct {
	Id       primitive.ObjectID `bson:"_id"`
	TypeId   string             `json:"typeid"`
	TypeName string             `json:"typename"`
}
