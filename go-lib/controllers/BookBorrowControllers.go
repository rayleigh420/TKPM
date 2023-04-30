package controllers

import (
	"context"
	// "fmt"
	// "net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		obj := bson.M{}
		book_rent_id := c.Param("book_rent_id")
		BookRentCollection.FindOne(ctx,bson.M{"book_rent_id":book_rent_id}).Decode(&obj)

		//create hire model
		bookBorrowModel := models.BookBorrowModel{}
		id,_ := gonanoid.Generate("123456789abcdef",12)
		bookBorrowModel.Id,_ = primitive.ObjectIDFromHex(id)
		bookBorrowModel.Book_id = obj["book_id"].(string)
		bookBorrowModel.User_id = obj["user_id"].(string)
	}
}

