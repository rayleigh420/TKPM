package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/baguette/go-lib/database"
	"github.com/baguette/go-lib/helper"
	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection = database.OpenCollection(database.Client, "users")
var NanoidString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func HashPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func VerifyPassword(userPassword string, foundUserPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "Password incorrect"
		check = false
	}
	return check, msg
}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		userModel := models.UserModel{}
		defer cancel()
		if err := c.BindJSON(&userModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad user model"})
			return
		}
		count, _ := UserCollection.CountDocuments(ctx, bson.D{
			{Key:"$or",Value: []bson.D{
				{{Key: "email", Value: userModel.Email}},
				{{Key: "phone", Value: userModel.Phone}},
			}},
		})
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone already exists"})
			return
		}
		userModel.Id = primitive.NewObjectID()
		userModel.User_id,_ = gonanoid.Generate(NanoidString,12)
		now := time.Now().Format(time.RFC3339)
		userModel.Created_at, _ = time.Parse(time.RFC3339, now)
		userModel.Updated_at, _ = time.Parse(time.RFC3339, now)
		userModel.Password, _ = HashPassword(userModel.Password)
		userModel.Role = "user"
		insertRes, insertErr := UserCollection.InsertOne(ctx, userModel)
		token,_ := helper.GenerateToken(userModel.User_id,userModel.Email,userModel.Role)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "insert user error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success":insertRes,
			"name":userModel.Name,
			"email":userModel.Email,
			"user_id":userModel.User_id,
			"role":userModel.Role,
			"token":token,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		userModel := models.UserModel{}
		foundUser := models.UserModel{}
		defer cancel()
		if err := c.ShouldBindJSON(&userModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad user model"})
			return
		}
		if err := UserCollection.FindOne(ctx, bson.D{{Key: "email", Value: userModel.Email}}).Decode(&foundUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email not found"})
			return
		}
		passwordIsValid, msg := VerifyPassword(userModel.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		token,_ := helper.GenerateToken(foundUser.User_id,foundUser.Email,foundUser.Role)
		//true
		c.JSON(http.StatusOK, gin.H{
			"status": "accepted",
			"name":foundUser.Name,
			"email":foundUser.Email,
			"user_id":foundUser.User_id,
			"role":foundUser.Role,
			"token":token,
		})
	}
}
