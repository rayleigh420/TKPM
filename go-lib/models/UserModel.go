package models

import (
	"time"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id				primitive.ObjectID		`bson:"_id"`
	Name			string					`json:"name"`
	Email			string					`json:"email"`
	Phone			string					`json:"phone"`
	Password        string					`json:"password"`	
	User_id			string					`json:"user_id"`
	Role			string					`json:"role"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at 		time.Time				`json:"updated_at"`
}