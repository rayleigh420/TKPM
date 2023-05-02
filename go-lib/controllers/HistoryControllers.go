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

func GetHistory() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.Background(),50*time.Second)
		defer cancel()
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 10
		startIndex := (page - 1) * recordPerPage
		skipStage := bson.D{{"$skip",startIndex}}
		lookupStage := bson.D{
			{"$lookup",bson.D{
				{"from","books"},
				{"localField","book_id"},
				{"foreignField","book_id"},
				{"as","book"},
			}},
		}
		// unsetStage := bson.D{{"$unset",bson.A{"_id"}}}
		unwindStage := bson.D{
			{"$unwind",bson.D{
				{"path","$book"},
				{"preserveNullAndEmptyArrays",false},
			}},
		}
		projectStage := bson.D{
			{"$project",bson.D{
				{"book.book_id",1},
				{"book.book_img",1},
				{"date_borrowed",1},
				{"date_return",1},
				{"user_id",1},
				{"status",1},
			}},
		}
		limitStage := bson.D{{"$limit",recordPerPage}}
		cursor,err := HistoryCollection.Aggregate(ctx,mongo.Pipeline{
			skipStage,limitStage,lookupStage,unwindStage,projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx,&result)
		c.JSON(http.StatusOK,result)
	}
}

func GetHistoryByUserId() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(context.Background(),50*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		matchStage := bson.D{{"$match",bson.D{{"user_id",user_id}}}}
		lookupStage := bson.D{
			{"$lookup",bson.D{
				{"from","books"},
				{"localField","book_id"},
				{"foreignField","book_id"},
				{"as","book"},
			}},
		}
		unwindStage := bson.D{
			{"$unwind",bson.D{
				{"path","$book"},
				{"preserveNullAndEmptyArrays",false},
			}},
		}
		projectStage := bson.D{
			{"$project",bson.D{
				{"book.book_id",1},
				{"book.book_img",1},
				{"date_borrowed",1},
				{"date_return",1},
				{"status",1},
			}},
		}
		cursor,err := HistoryCollection.Aggregate(ctx,mongo.Pipeline{
			matchStage,lookupStage,unwindStage,projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx,&result)
		c.JSON(http.StatusOK,result)
	}
}