package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateRentRequest() error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	yest, _ := time.Parse(time.RFC3339, time.Now().Add(-24*time.Hour).Format(time.RFC3339))
	//delete rent request and update status to ready if over 24 hrs
	updateObj := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "ready"}, {Key: "updated_at", Value: now}}}}

	//update book set amount to normal
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{
					{Key: "updated_at", Value: bson.D{
						{Key: "$lte", Value: yest},
					}},
					{Key: "status", Value: "booked"},
				}},
			},
		}},
	}
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$book_id"},
			{Key: "count", Value: bson.M{"$sum": 1}},
		}},
	}
	bookCursor, _ := BookDetailCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage,
	})
	res := []bson.M{}
	bookCursor.All(ctx, &res)
	updateModels := []mongo.WriteModel{}
	for _, amount := range res {
		filter := bson.M{"book_id": amount["_id"].(string)}
		update := bson.M{"$inc": bson.M{"amount": amount["count"].(int32)}}
		model := mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update)
		updateModels = append(updateModels, model)
	}
	//done step update
	if len(updateModels) != 0 {
		if _, bulkErr := BookCollection.BulkWrite(ctx, updateModels); bulkErr != nil {
			return bulkErr
		}
	}
	//step update bookdetail
	BookDetailCollection.UpdateMany(ctx, bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{
				{Key: "updated_at", Value: bson.D{
					{Key: "$lte", Value: yest},
				}},
				{Key: "status", Value: "booked"},
			}}}}, updateObj)
	//step delete rent request
	BookRentCollection.DeleteMany(ctx, bson.D{
		{Key: "reserve_date", Value: bson.D{
			{Key: "$lte", Value: yest},
		}},
	})
	return nil
}
func GetRentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		// now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// yest, _ := time.Parse(time.RFC3339, time.Now().Add(-24*time.Hour).Format(time.RFC3339))
		// //delete rent request and update status to ready if over 24 hrs
		// updateObj := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "ready"}, {Key: "updated_at", Value: now}}}}

		// //update book set amount to normal
		// matchStage := bson.D{
		// 	{Key: "$match", Value: bson.D{
		// 		{Key: "$and", Value: bson.A{
		// 			bson.D{
		// 				{Key: "updated_at", Value: bson.D{
		// 					{Key: "$lte", Value: yest},
		// 				}},
		// 				{Key: "status", Value: "booked"},
		// 			}},
		// 		},
		// 	}},
		// }
		// groupStage := bson.D{
		// 	{Key: "$group", Value: bson.D{
		// 		{Key: "_id", Value: "$book_id"},
		// 		{Key: "count", Value: bson.M{"$sum": 1}},
		// 	}},
		// }
		// bookCursor, _ := BookDetailCollection.Aggregate(ctx, mongo.Pipeline{
		// 	matchStage, groupStage,
		// })
		// res := []bson.M{}
		// bookCursor.All(ctx, &res)
		// updateModels := []mongo.WriteModel{}
		// for _, amount := range res {
		// 	filter := bson.M{"book_id": amount["_id"].(string)}
		// 	update := bson.M{"$inc": bson.M{"amount": amount["count"].(string)}}
		// 	model := mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update)
		// 	updateModels = append(updateModels, model)
		// }
		// //done step update
		// BookCollection.BulkWrite(ctx, updateModels)
		// //step update bookdetail
		// BookDetailCollection.UpdateMany(ctx, bson.D{
		// 	{Key: "$and", Value: bson.A{
		// 		bson.D{
		// 			{Key: "updated_at", Value: bson.D{
		// 				{Key: "$lte", Value: yest},
		// 			}},
		// 			{Key: "status", Value: "booked"},
		// 		}}}}, updateObj)
		// //step delete rent request
		// BookRentCollection.DeleteMany(ctx, bson.D{
		// 	{Key: "reserve_date", Value: bson.D{
		// 		{Key: "$lte", Value: yest},
		// 	}},
		// })
		t1 := time.Now()
		if err := UpdateRentRequest(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		t2 := time.Now()
		result := []bson.M{}
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
		rentListCursor, rentListErr := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage, unwindStage, projectStage,
		})
		if rentListErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list"})
			return
		}
		rentListCursor.All(ctx, &result)
		c.JSON(http.StatusOK, gin.H{
			"res":         result,
			"time update": t2.Sub(t1),
		})
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
func AbortRentRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book_rent_id := c.Param("book_rent_id")
		bookRentModel := bson.M{}
		BookRentCollection.FindOneAndDelete(ctx, bson.M{"book_rent_id": book_rent_id}).Decode(&bookRentModel)
		book_detail_id := bookRentModel["book_detail_id"].(string)
		book_id := bookRentModel["book_id"].(string)
		updateObj := bson.M{"$set": bson.M{"status": "ready", "updated_at": now}}
		updateObj2 := bson.M{"$inc": bson.M{"amount": 1}}
		BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": book_detail_id}, updateObj)
		if _, err := BookCollection.UpdateOne(ctx, bson.M{"book_id": book_id}, updateObj2); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":         "abort succeeded",
			"book_detail_id": book_detail_id,
			"book_id":        book_id,
		})
	}
}
