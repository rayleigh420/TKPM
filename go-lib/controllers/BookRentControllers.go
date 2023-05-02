package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetRentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		// now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		yest, _ := time.Parse(time.RFC3339, time.Now().Add(-24 * time.Hour).Format(time.RFC3339))
		//delete rent request and update status to ready if over 24 hrs
		updateObj := bson.D{{"$set", bson.D{{"status", "ready"}}}}
		// del := options.DeleteOptions{}
		// BookDetailCollection.UpdateMany(ctx, bson.D{
		// 	{"reserve_date", bson.D{{"gt", now}}},
		// }, updateObj)
		// BookRentCollection.DeleteMany(ctx, bson.D{
		// 	{Key: "reserve_date", Value: bson.D{{Key: "gt", Value: now}}},
		// })
		res := []bson.M{}
		cursor, _ := BookRentCollection.Find(ctx, bson.D{
			{Key: "reserve_date", Value: bson.D{
				{Key: "$lte", Value: yest},
			}},
		})
		cursor.All(ctx, &res)
		book_detail := []string{}
		for _,val := range res {
			book_detail = append(book_detail, val["book_detail_id"].(string))
		}
		BookDetailCollection.UpdateMany(ctx,bson.D{
			{Key: "book_detail_id",Value: bson.D{
				{Key: "$in",Value: book_detail},
			}},
		},updateObj)
		BookRentCollection.DeleteMany(ctx,bson.D{
			{Key: "book_detail_id",Value: bson.D{
				{Key: "$in",Value: book_detail},
			}},
		})
		
		c.JSON(http.StatusOK, res)
		// result := []bson.M{}
		// lookupStage := bson.D{
		// 	{"$lookup", bson.D{
		// 		{"from", "books"},
		// 		{"localField", "book_id"},
		// 		{"foreignField", "book_id"},
		// 		{"as", "book"},
		// 	}},
		// }
		// unwindStage := bson.D{{"$unwind", bson.D{{"path", "$book"}, {"preserveNullAndEmptyArrays", false}}}}
		// projectStage := bson.D{
		// 	{"$project", bson.D{
		// 		{"_id", 0},
		// 		{"book._id", 0},
		// 	}},
		// }
		// rentListCursor, rentListErr := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
		// 	lookupStage, unwindStage, projectStage,
		// })
		// if rentListErr != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list"})
		// 	return
		// }
		// rentListCursor.All(ctx, &result)
		// c.JSON(http.StatusOK, result)
	}
}

func GetRentListOfBook(book_id string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	// now, _ := time.Parse(time.RFC3339, time.Now().Add(-24*time.Hour).Format(time.RFC3339))
	// updateObj := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "ready"}}}}
	// BookDetailCollection.UpdateMany(ctx, bson.D{
	// 	{Key: "reserve_date", Value: bson.D{{Key: "gt", Value: now}}},
	// }, updateObj)
	// BookRentCollection.DeleteMany(ctx, bson.D{
	// 	{Key: "reserve_date", Value: bson.D{{Key: "gt", Value: now}}},
	// })
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{{Key: "book_id", Value: book_id}}},
	}
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "books"},
			{Key: "localField", Value: "book_id"},
			{Key: "foreignField", Value: "book_id"},
			{Key: "as", Value: "book"},
		}},
	}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "book._id", Value: 0},
		}},
	}
	result := []bson.M{}
	rentListCursor, rentListErr := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, projectStage,
	})
	if rentListErr != nil {
		return result, rentListErr
	}
	rentListCursor.All(ctx, &result)
	return result, nil
}

func GetRentListById() gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id := c.Param("book_id")
		result, err := GetRentListOfBook(book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error happened",
			})
		}
		c.JSON(http.StatusOK, result)
	}
}
