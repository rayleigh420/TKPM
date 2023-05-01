package controllers

import (
	"context"
	"net/http"

	// "fmt"
	// "net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/matoous/go-nanoid/v2"
)

func HireABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		//rent_id
		//user_id
		//book_detail_id
		now,_ := time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		month1,_ := time.Parse(time.RFC3339,time.Now().Add(30*24*time.Hour).Format(time.RFC3339))
		obj := bson.M{}
		book_rent_id := c.Param("book_rent_id")
		BookRentCollection.FindOneAndDelete(ctx,bson.M{"book_rent_id":book_rent_id}).Decode(&obj)

		//create borrow model
		bookBorrowModel := models.BookBorrowModel{}
		bookBorrowModel.Id = primitive.NewObjectID()
		bookBorrowModel.Book_id = obj["book_id"].(string)
		bookBorrowModel.User_id = obj["user_id"].(string)
		bookReady,err := FindBookToRentWithId(bookBorrowModel.Book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"cannot find book"})
			return
		}
		bookBorrowModel.Book_detail_id = bookReady["book_detail_id"].(string)
		bookBorrowModel.Book_hire_id,_ = gonanoid.Generate(NanoidString,12)
		bookBorrowModel.Date_borrowed = now
		bookBorrowModel.Date_end = month1
		
		//update book detail status
		updateObj := bson.D{
			{Key:"$set",Value:bson.D{{Key:"status",Value:"borrowed"}}},
		}
		
		
		//create history borrowing
		historyModel := models.HistoryModel{}
		historyModel.Id = primitive.NewObjectID()
		historyModel.Date_borrowed = now
		historyModel.History_id,_ = gonanoid.Generate(NanoidString,12)
		historyModel.Book_detail_id = bookBorrowModel.Book_detail_id
		historyModel.Status = "borrowing"
		
		BookBorrowedCollection.InsertOne(ctx,bookBorrowModel)
		BookDetailCollection.UpdateOne(ctx,bson.M{"book_detail_id":bookBorrowModel.Book_detail_id},updateObj)
		HistoryCollection.InsertOne(ctx,historyModel)
	}
}

func GetBorrowList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		borrowList := []bson.M{}
		lookupStage := bson.D{
			{Key: "$lookup",Value: bson.D{
				{Key: "from",Value: "users"},
				{Key: "localField",Value: "user_id"},
				{Key: "foreignField",Value: "user_id"},
				{Key: "as",Value: "user"},
			}},
		}
		projectStage := bson.D{
			{Key: "$project",Value: bson.D{
				{Key: "_id",Value: 0},
				{Key: "user._id",Value: 0},
				{Key: "user.password",Value: 0},
				{Key: "user.email",Value: 0},
				{Key: "user.phone",Value: 0},
				{Key: "user.role",Value: 0},
				{Key: "user.created_at",Value: 0},
				{Key: "user.updated_at",Value: 0},
				{Key:"user_id",Value:0},
			}},
		}
		cursor,_:=BookBorrowedCollection.Aggregate(ctx,mongo.Pipeline{
			lookupStage,projectStage,
		})
		cursor.All(ctx,&borrowList)
		c.JSON(http.StatusOK,borrowList)
	}
}
