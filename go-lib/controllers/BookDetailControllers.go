package controllers

import (
	"bytes"
	"context"
	"strings"

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
	"go.mongodb.org/mongo-driver/mongo/options"
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
		opt := options.FindOne().SetSort(bson.M{"$natural": -1})
		res := models.BookDetailModel{}
		BookDetailCollection.FindOne(ctx, bson.M{"book_id":book_id}, opt).Decode(&res)
		t := strings.Split(res.Book_id, "B")[1]
		splits := bytes.Buffer{}
		splits.WriteString("D")
		splits.WriteString(t)
		t2 := strings.Split(res.Book_detail_id,splits.String())[1] 
		count, _ := strconv.Atoi(t2)
		count++;
		counts := strconv.Itoa(count)
		id := bytes.Buffer{}
		id.WriteString(splits.String())
		id.WriteString(counts)
		// id = "D" + t + counts
		bookDetailModel.Id = primitive.NewObjectID()
		bookDetailModel.Book_id = book_id
		bookDetailModel.Book_detail_id = id.String()
		bookDetailModel.Status = "ready"
		bookDetailModel.Created_at = now
		bookDetailModel.Updated_at = now

		updateObj := bson.M{"$inc": bson.M{"amount": 1},"$set":bson.M{"updated_at":now}}
		_,err := BookCollection.UpdateOne(ctx,bson.M{"book_id":book_id},updateObj)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error updating"})
			return
		}
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

// func UpdateBookDetail() gin.HandlerFunc{
// 	return func (c *gin.Context)  {
// 		ctx,cancel := context.WithTimeout(context.TODO(),50*time.Second)
// 		defer cancel()

// 	}
// }

func DeleteBookVersion() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.TODO(),50*time.Second)
		defer cancel()
		id := c.Param("book_detail_id");
		bookDetailModel := models.BookDetailModel{}
		if err := BookDetailCollection.FindOneAndDelete(ctx,bson.M{"book_detail_id":id}).Decode(&bookDetailModel);err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error deleting book version"})
			return
		}
		updateObj := bson.M{"$inc": bson.M{"amount": -1}}
		BookCollection.UpdateOne(ctx,bson.M{"book_id":bookDetailModel.Book_id},updateObj)
		BookRentCollection.DeleteMany(ctx,bson.M{"book_detail_id":id})
		BookBorrowedCollection.DeleteMany(ctx,bson.M{"book_detail_id":id})
		HistoryCollection.DeleteMany(ctx,bson.M{"book_detail_id":id})
		
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
			"book_detail_id":id,
		})
	}
}