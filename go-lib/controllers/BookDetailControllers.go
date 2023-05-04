package controllers

import (
	"context"
	// "fmt"
	"strconv"

	// "fmt"
	"net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/matoous/go-nanoid/v2"
)

func CreateBooKDetail() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.TODO(),50*time.Second)
		defer cancel()
		now,_ := time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		book_id := c.Param("book_id")
		bookDetailModel := models.BookDetailModel{}
		if err := c.ShouldBindJSON(&bookDetailModel);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":"bad model"})
			return
		}
		book := models.BookModel{}
		// id,_ := gonanoid.Generate("1234567890abcdef",12)
		BookCollection.FindOne(ctx,bson.M{"book_id":book_id}).Decode(&book)
		if len(book.Book_id) == 0 {
			c.JSON(http.StatusBadRequest,gin.H{"error":"can't find book"})
			return
		}
		count,_ := BookDetailCollection.CountDocuments(ctx,bson.M{"book_id":book_id})
		count++
		counts := strconv.Itoa(int(count))
		id := "D" + book_id[1:] + counts
		bookDetailModel.Id = primitive.NewObjectID()
		bookDetailModel.Book_id = book_id
		bookDetailModel.Book_detail_id = id
		bookDetailModel.Status = "ready"
		bookDetailModel.Created_at = now
		bookDetailModel.Updated_at = now

		updateObj := bson.M{"$inc": bson.M{"amount": 1},"updated_at":now}
		BookCollection.UpdateOne(ctx,bson.M{"book_id":book_id},updateObj)
		_,insertErr := BookDetailCollection.InsertOne(ctx,bookDetailModel)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":insertErr})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
			"book_detail_id":bookDetailModel.Book_detail_id,
		})
	}
}