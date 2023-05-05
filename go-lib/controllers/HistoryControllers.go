package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	// "github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 10
		startIndex := (page - 1) * recordPerPage
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "localField", Value: "book_id"},
				{Key: "foreignField", Value: "book_id"},
				{Key: "as", Value: "book"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "user"},
			}},
		}
		// unsetStage := bson.D{{"$unset",bson.A{"_id"}}}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$book"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$user"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "book.book_id", Value: 1},
				{Key: "book.book_img", Value: 1},
				{Key: "user.name", Value: 1},
				{Key: "user.avatar", Value: 1},
				{Key: "date_borrowed", Value: 1},
				{Key: "date_return", Value: 1},
				{Key: "user_id", Value: 1},
				{Key: "status", Value: 1},
			}},
		}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			skipStage, limitStage, lookupStage, lookupStage2, unwindStage, unwindStage2, projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx, &result)
		c.JSON(http.StatusOK, result)
	}
}

func GetHistoryByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "localField", Value: "book_id"},
				{Key: "foreignField", Value: "book_id"},
				{Key: "as", Value: "book"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$book"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "book.book_id", Value: 1},
				{Key: "book.book_img", Value: 1},
				{Key: "date_borrowed", Value: 1},
				{Key: "date_return", Value: 1},
				{Key: "status", Value: 1},
			}},
		}
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, unwindStage, projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx, &result)
		c.JSON(http.StatusOK, result)
	}
}

func GetReturnedBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		res := []bson.M{}
		// cursor,_ := HistoryCollection.Find(ctx,bson.M{"status":"borrowing"})
		// cursor.All(ctx,&res)
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "status", Value: "returned"},
			}},
		}
		sortStage := bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "date_borrowed", Value: -1},
			}},
		}
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
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: false}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: false}}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "history_id", Value: 1},
				{Key: "user.avatar", Value: 1},
				{Key: "user.name", Value: 1},
				{Key: "book.book_id", Value: 1},
				{Key: "book.name", Value: 1},
				{Key: "book.book_img", Value: 1},
				{Key: "date_borrowed", Value: 1},
				{Key: "date_return", Value: 1},
			}},
		}
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			sortStage, matchStage, lookupStage, lookupStage2, unwindStage, unwindStage2, projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		cursor.All(ctx, &res)
		c.JSON(http.StatusOK, res)
	}
}
