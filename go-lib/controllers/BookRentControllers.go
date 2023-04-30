package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		now, _ := time.Parse(time.RFC3339, time.Now().Add(-24*time.Hour).Format(time.RFC3339))
		//delete rent request that is over 24 hrs
		BookRentCollection.DeleteMany(ctx, bson.D{
			{"date_reserved", bson.D{{"gt", now}}},
		})
		result := []bson.M{}
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "books"},
				{"localField", "book_id"},
				{"foreignField", "book_id"},
				{"as", "book"},
			}},
		}
		unwindStage := bson.D{{"$unwind", bson.D{{"path", "$book"}, {"preserveNullAndEmptyArrays", false}}}}
		rentListCursor,rentListErr := BookRentCollection.Aggregate(ctx,mongo.Pipeline{
			lookupStage,unwindStage,
		})
		if rentListErr != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error getting list"})
			return
		}
		rentListCursor.All(ctx, &result)
		c.JSON(http.StatusOK, result)
	}
}
