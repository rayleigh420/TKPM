package controllers

import (
	"context"
	// "fmt"
	"net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/matoous/go-nanoid/v2"
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
		id,_ := gonanoid.Generate("1234567890abcdef",12)
		bookDetailModel.Book_id = book_id
		bookDetailModel.Book_detail_id = id
		bookDetailModel.Status = "ready"
		bookDetailModel.Created_at = now
		bookDetailModel.Updated_at = now
		insertRes,_ := BookDetailCollection.InsertOne(ctx,bookDetailModel)
		c.JSON(http.StatusOK,insertRes)
	}
}