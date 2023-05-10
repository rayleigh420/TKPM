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
	// "go.mongodb.org/mongo-driver/mongo"
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
		// if emailErr := helper.CheckEmail(userModel.Email);emailErr != nil {

		// }
		count, _ := UserCollection.CountDocuments(ctx, bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "email", Value: userModel.Email}},
				{{Key: "phone", Value: userModel.Phone}},
			}},
		})
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone already exists"})
			return
		}
		userModel.Id = primitive.NewObjectID()
		userModel.User_id, _ = gonanoid.Generate(NanoidString, 12)
		now := time.Now().Format(time.RFC3339)
		userModel.Created_at, _ = time.Parse(time.RFC3339, now)
		userModel.Updated_at, _ = time.Parse(time.RFC3339, now)
		userModel.Password, _ = HashPassword(userModel.Password)
		userModel.Role = "user"
		insertRes, insertErr := UserCollection.InsertOne(ctx, userModel)
		token, _ := helper.GenerateToken(userModel.Name, userModel.Email, userModel.Role, userModel.User_id, userModel.Avatar)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "insert user error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": insertRes,
			"name":    userModel.Name,
			"email":   userModel.Email,
			"user_id": userModel.User_id,
			"avatar":  userModel.Avatar,
			"role":    userModel.Role,
			"token":   token,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		userModel := bson.M{}
		foundUser := models.UserModel{}
		defer cancel()
		// t1 := time.Now()
		if err := c.ShouldBindJSON(&userModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad user model"})
			return
		}
		// t2 := time.Now()
		if err := UserCollection.FindOne(ctx, bson.D{{Key: "email", Value: userModel["email"].(string)}}).Decode(&foundUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email not found"})
			return
		}
		// t3 := time.Now()
		passwordIsValid, msg := VerifyPassword(userModel["password"].(string), foundUser.Password)
		// t4 := time.Now()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		token, _ := helper.GenerateToken(foundUser.Name, foundUser.Email, foundUser.Role, foundUser.User_id, foundUser.Avatar)
		// t5 := time.Now()
		//true
		c.JSON(http.StatusOK, gin.H{
			"status":  "accepted",
			"name":    foundUser.Name,
			"email":   foundUser.Email,
			"user_id": foundUser.User_id,
			"avatar":  foundUser.Avatar,
			"role":    foundUser.Role,
			"token":   token,
		})
	}
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		claim, err := helper.ValidateToken(token)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		res := bson.M{}
		res["name"] = claim.Name
		res["email"] = claim.Email
		res["token"] = token
		res["role"] = claim.Role
		res["user_id"] = claim.User_id
		res["avatar"] = claim.Avatar
		res["exp"] = claim.ExpiresAt
		c.JSON(http.StatusOK, res)
	}
}

func Updateuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		userModel := models.UserModel{}
		if err := c.Bind(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		updateObj := bson.D{}
		if userModel.Email != "" {
			count, _ := UserCollection.CountDocuments(ctx, bson.D{{Key: "email", Value: userModel.Email}})
			if count > 0 {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "email", Value: userModel.Email})
		}
		if userModel.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: userModel.Name})
		}
		if userModel.Avatar != "" {
			updateObj = append(updateObj, bson.E{Key: "avatar", Value: userModel.Avatar})
		}
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: now})
		UserCollection.UpdateOne(ctx, bson.M{"user_id": user_id}, bson.D{
			{Key: "$set", Value: updateObj},
		})
		user := models.UserModel{}
		UserCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user)
		token, err := helper.GenerateToken(user.Name, user.Email, user.Role, user.User_id, user.Avatar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		// rec := bson.M{}
		// c.ShouldBindJSON(&rec)
		// token := rec["token"].(string)
		// claim, msg := helper.ValidateToken(token)
		// if msg != "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		// 	return
		// }
		// if claim.Role != "admin" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		// 	return
		// }
		users := []bson.M{}
		cursor, cursorErr := UserCollection.Find(ctx, bson.M{})
		if cursorErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": cursorErr.Error()})
			return
		}
		if err := cursor.All(ctx, &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		_, deleteRes := UserCollection.DeleteOne(ctx, bson.M{"user_id": user_id})
		if deleteRes != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"user_id": user_id,
		})
	}
}
