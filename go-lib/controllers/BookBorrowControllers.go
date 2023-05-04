package controllers

import (
	"context"
	"net/http"

	// "fmt"
	// "net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/helper"
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
		//token
		rec := bson.M{}
		c.ShouldBindJSON(&rec)
		book_rent_id := rec["book_rent_id"].(string)
		token := rec["token"].(string)
		claim, msg := helper.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		if claim.Role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
			return
		}
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		month1, _ := time.Parse(time.RFC3339, time.Now().Add(30*24*time.Hour).Format(time.RFC3339))
		obj := bson.M{}
		if err := BookRentCollection.FindOne(ctx, bson.M{"book_rent_id": book_rent_id}).Decode(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "what up"})
			return
		}
		//create borrow model
		bookBorrowModel := models.BookBorrowModel{}
		bookBorrowModel.Id = primitive.NewObjectID()
		bookBorrowModel.Book_id = obj["book_id"].(string)
		bookBorrowModel.User_id = obj["user_id"].(string)
		bookBorrowModel.Book_detail_id = obj["book_detail_id"].(string)
		bookBorrowModel.Book_hire_id, _ = gonanoid.Generate(NanoidString, 12)
		bookBorrowModel.Date_borrowed = now
		bookBorrowModel.Date_end = month1

		//update book detail status
		updateObj := bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: "borrowed"}, {Key: "updated_at", Value: now}}},
		}

		//create history borrowing
		historyModel := models.HistoryModel{}
		historyModel.Id = primitive.NewObjectID()
		historyModel.Book_id = obj["book_id"].(string)
		historyModel.User_id = obj["user_id"].(string)
		historyModel.Book_detail_id = obj["book_detail_id"].(string)
		historyModel.Date_borrowed = now
		historyModel.History_id, _ = gonanoid.Generate(NanoidString, 12)
		historyModel.Book_hire_id = bookBorrowModel.Book_hire_id
		historyModel.Status = "borrowing"
		BookRentCollection.DeleteOne(ctx, bson.M{"book_rent_id": obj["book_rent_id"].(string)})
		BookBorrowedCollection.InsertOne(ctx, bookBorrowModel)
		BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookBorrowModel.Book_detail_id}, updateObj)
		HistoryCollection.InsertOne(ctx, historyModel)
		c.JSON(http.StatusOK, gin.H{
			"status":         "success",
			"history_id":     historyModel.History_id,
			"borrowed_id":    bookBorrowModel.Book_hire_id,
			"book_detail_id": bookBorrowModel.Book_detail_id,
			"book_id":        bookBorrowModel.Book_id,
			"date_borrowed":  bookBorrowModel.Date_borrowed,
			"book_hire_id": bookBorrowModel.Book_hire_id,
		})
	}
}



func GetBorrowList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		borrowList := []bson.M{}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "localField", Value: "book_id"},
				{Key: "foreignField", Value: "book_id"},
				{Key: "as", Value: "book"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"},{Key: "preserveNullAndEmptyArrays", Value: false}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"},{Key: "preserveNullAndEmptyArrays", Value: false}}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "book_hire_id",Value: 1},
				{Key: "book_detail_id",Value: 1},
				{Key: "user.avatar",Value: 1},
				{Key: "user.name",Value: 1},
				{Key: "book.book_id",Value: 1},
				{Key: "book.name",Value: 1},
				{Key: "date_borrowed",Value: 1},
				{Key: "date_end",Value: 1},
			}},
		}
		cursor, _ := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage,lookupStage2,unwindStage,unwindStage2, projectStage,
		})
		cursor.All(ctx, &borrowList)
		c.JSON(http.StatusOK, borrowList)
	}
}

func GetBorrowListOfABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		borrowList := []bson.M{}
		book_id := c.Param("book_id")
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "book_id", Value: book_id}}}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "localField", Value: "book_id"},
				{Key: "foreignField", Value: "book_id"},
				{Key: "as", Value: "book"},
			}},
		}
		// projectStage := bson.D{
		// 	{Key: "$project",Value: bson.D{
		// 		{Key: "_id",Value: 0},
		// 		{Key: "user._id",Value: 0},
		// 		{Key: "book._id",Value: 0},
		// 		{Key: "user.password",Value: 0},
		// 		{Key: "user.email",Value: 0},
		// 		{Key: "user.phone",Value: 0},
		// 		{Key: "user.role",Value: 0},
		// 		{Key: "user.created_at",Value: 0},
		// 		{Key: "user.updated_at",Value: 0},
		// 		{Key:"user_id",Value:0},
		// 	}},
		// }
		unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
		unwindStage2 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				// {Key: "book.book_id",Value: 1},
				// {Key: "book.book_img",Value: 1},
				// {Key: "book.name",Value: 1},
				// {Key: "book.book_detail_id",Value: 1},
				// {Key: "book.book_hire_id",Value: 1},
				// {Key: "book_id",Value: 1},
				// {Key: "book_id",Value: 1},
				// {Key: "date_borrowed",Value: 1},
				// {Key: "date_end",Value: 1},
				// {Key: "user.user_id",Value: 1},
				// {Key: "user.name",Value: 1},
			}},
		}
		cursor, _ := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2, unwindStage, unwindStage2, projectStage,
		})
		cursor.All(ctx, &borrowList)
		c.JSON(http.StatusOK, borrowList)
	}
}

func ReturnABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.Background(),50*time.Second)
		defer cancel()
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		rec := bson.M{}
		c.ShouldBindJSON(&rec)
		token := rec["token"].(string)
		claim,msg := helper.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest,gin.H{"error":msg})
			return
		}
		if claim.Role != "admin"{
			c.JSON(http.StatusBadRequest,gin.H{"error":"unauthorized"})
			return
		}
		book_hire_id := c.Param("book_hire_id")
		bookBorrowModel := bson.M{}
		BookBorrowedCollection.FindOneAndDelete(ctx,bson.M{"book_hire_id":book_hire_id}).Decode(&bookBorrowModel)
		updateObj := bson.M{"$set": bson.M{"status": "ready", "updated_at": now}}
		updateHis := bson.M{"$set": bson.M{"status": "returned", "date_return": now,"book_hire_id":""}}
		updateObj2 := bson.M{"$inc": bson.M{"amount": 1,"borrowed_quantity":1}}
		_,err1 := BookDetailCollection.UpdateOne(ctx,bson.M{"book_detail_id":bookBorrowModel["book_detail_id"].(string)},updateObj)
		_,err2 := HistoryCollection.UpdateOne(ctx,bson.M{"book_hire_id":bookBorrowModel["book_hire_id"].(string)},updateHis)
		_,err3 := BookCollection.UpdateOne(ctx,bson.M{"book_id":bookBorrowModel["book_id"].(string)},updateObj2)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error 1"})
			return
		}
		if err2 != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error 2"})
			return
		}
		if err3 != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error 3"})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"status":"returned",
		})
	}
}