package controllers

import (
	"context"
	"fmt"
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
	// min1, _ := time.Parse(time.RFC3339, time.Now().Add(-1*time.Minute).Format(time.RFC3339))
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
	// matchStage := bson.D{
	// 	{Key: "$match", Value: bson.D{
	// 		{Key: "$and", Value: bson.A{
	// 			bson.D{
	// 				{Key: "updated_at", Value: bson.D{
	// 					{Key: "$lte", Value: min1},
	// 				}},
	// 				{Key: "status", Value: "booked"},
	// 			}},
	// 		},
	// 	}},
	// }
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
	// BookDetailCollection.UpdateMany(ctx, bson.D{
	// 	{Key: "$and", Value: bson.A{
	// 		bson.D{
	// 			{Key: "updated_at", Value: bson.D{
	// 				{Key: "$lte", Value: min1},
	// 			}},
	// 			{Key: "status", Value: "booked"},
	// 		}}}}, updateObj)
	//step delete rent request
	BookRentCollection.DeleteMany(ctx, bson.D{
		{Key: "reserve_date", Value: bson.D{
			{Key: "$lte", Value: yest},
		}},
	})
	// BookRentCollection.DeleteMany(ctx, bson.D{
	// 	{Key: "reserve_date", Value: bson.D{
	// 		{Key: "$lte", Value: min1},
	// 	}},
	// })
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
		// t1 := time.Now()
		if err := UpdateRentRequest(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		// t2 := time.Now()
		result := []bson.M{}
		// lookupStage := bson.D{
		// 	{Key: "$lookup", Value: bson.D{
		// 		{Key: "from", Value: "books"},
		// 		{Key: "localField", Value: "book_id"},
		// 		{Key: "foreignField", Value: "book_id"},
		// 		{Key: "as", Value: "book"},
		// 	}},
		// }
		// lookupStage2 := bson.D{
		// 	{Key: "$lookup", Value: bson.D{
		// 		{Key: "from", Value: "users"},
		// 		{Key: "localField", Value: "user_id"},
		// 		{Key: "foreignField", Value: "user_id"},
		// 		{Key: "as", Value: "user"},
		// 	}},
		// }
		// lookupStage3 := bson.D{
		// 	{Key: "$lookup", Value: bson.D{
		// 		{Key: "from", Value: "book_detail"},
		// 		{Key: "localField", Value: "book_detail_id"},
		// 		{Key: "foreignField", Value: "book_detail_id"},
		// 		{Key: "as", Value: "book_detail"},
		// 	}},
		// }
		// unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}}}}
		// unwindStage2 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}}}}
		// unwindStage3 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}}}
		
		// projectStage := bson.D{
		// 	{Key: "$project", Value: bson.D{
		// 		{Key: "_id", Value: 0},
		// 		{Key: "book._id", Value: 0},
		// 		{Key: "user._id", Value: 0},
		// 		{Key: "book_detail._id", Value: 0},
		// 	}},
		// }
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage3 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_detail"},
				{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book_detail"},
			}},
		}
		lookupStage4 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
				{Key: "pipeline", Value: bson.A{
					// bson.D{{Key: "$unwind",Value: "$$book"}},
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
		}
		rentListCursor, rentListErr := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage,lookupStage2,lookupStage3, unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
		})
		if rentListErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list"})
			return
		}
		rentListCursor.All(ctx, &result)
		c.JSON(http.StatusOK, result)
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
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}}}}
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

func GetRentListOfBook2(book_id string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	addFieldStage := bson.D{
		{Key: "$addFields",Value: bson.D{{Key: "status",Value: "booked"}}},
	}
	matchStage := bson.D{
		{Key: "$match",Value: bson.D{{Key: "book_id",Value: book_id}}},
	}
	// lookupStage := bson.D{
	// 	{"$lookup",bson.D{
	// 		{"from","books"},
	// 		{"localField","book_id"},
	// 		{"foreignField","book_id"},
	// 		{"as","book"},
	// 	}},
	// }
	// lookupStage2 := bson.D{
	// 	{"$lookup",bson.D{
	// 		{"from","users"},
	// 		{"localField","user_id"},
	// 		{"foreignField","user_id"},
	// 		{"as","user"},
	// 	}},
	// }
	// unwindStage := bson.D{
	// 	{"$unwind",bson.D{{"path","$book"},{"preserveNullAndEmptyArrays",false}}},
	// }
	// unwindStage2 := bson.D{
	// 	{"$unwind",bson.D{{"path","$user"},{"preserveNullAndEmptyArrays",false}}},
	// }
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "books"},
			{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "book"},
		}},
	}
	lookupStage2 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "user"},
		}},
	}
	lookupStage3 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "book_detail"},
			{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "book_detail"},
		}},
	}
	lookupStage4 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "book_types"},
			{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
			{Key: "pipeline", Value: bson.A{
				// bson.D{{Key: "$unwind",Value: "$$book"}},
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "type"},
		}},
	}

	unwindStage := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
	}
	unwindStage2 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
	}
	unwindStage3 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
	}
	unwindStage4 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
	}
	cursor, err := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
		addFieldStage,matchStage, lookupStage, lookupStage2,lookupStage3,unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
	})
	if err != nil {
		return []bson.M{},err
	}
	result := []bson.M{}
	cursor.All(ctx, &result)
	return result,nil
}

func GetRentListById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		rent_id := c.Query("search_id")
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "book_rent_id", Value: rent_id},
			}},
		}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage3 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_detail"},
				{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book_detail"},
			}},
		}
		lookupStage4 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
				{Key: "pipeline", Value: bson.A{
					// bson.D{{Key: "$unwind",Value: "$$book"}},
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "type"},
			}},
		}

		unwindStage := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
		}
		// setStage := bson.D{
		// 	{"$set", bson.D{
		// 		{"book", "$book.0"},
		// 		{"user", "$user.0"},
		// 	}},
		// }
		// projectStage := bson.D{
		// 	{"$project", bson.D{
		// 		{"_id", 0},
		// 		{"user",0},
		// 		{"book",0},
		// 	}},
		// }
		cursor, err := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2,lookupStage3,unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx, &result)
		if len(result) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "request not exist"})
			return
		}
		c.JSON(http.StatusOK, result[0])
	}
}

func GetRequest(id string) (bson.M,error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		rent_id := id
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "book_rent_id", Value: rent_id},
			}},
		}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage3 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_detail"},
				{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book_detail"},
			}},
		}
		lookupStage4 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
				{Key: "pipeline", Value: bson.A{
					// bson.D{{Key: "$unwind",Value: "$$book"}},
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "type"},
			}},
		}

		unwindStage := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
		}
		cursor, err := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2,lookupStage3,unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
		})
		result := []bson.M{}
		if err != nil {
			return result[0],err
		}
		cursor.All(ctx, &result)
		if len(result) == 0 {
			err := fmt.Errorf("cant find request")
			return result[0],err
		}
		result[0]["status"] = "booked"
		return result[0],nil
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

func GetRentListOfUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "user_id", Value: user_id},
			}},
		}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "books"},
				{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "user"},
			}},
		}
		lookupStage3 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_detail"},
				{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "book_detail"},
			}},
		}
		lookupStage4 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
				{Key: "pipeline", Value: bson.A{
					// bson.D{{Key: "$unwind",Value: "$$book"}},
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
					bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
				}},
				{Key: "as", Value: "type"},
			}},
		}

		unwindStage := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
		}
		// setStage := bson.D{
		// 	{"$set", bson.D{
		// 		{"book", "$book.0"},
		// 		{"user", "$user.0"},
		// 	}},
		// }
		// projectStage := bson.D{
		// 	{"$project", bson.D{
		// 		{"_id", 0},
		// 		{"user",0},
		// 		{"book",0},
		// 	}},
		// }
		cursor, err := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2,lookupStage3,unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx, &result)
		// if len(result) == 0 {
		// 	c.JSON(http.StatusOK, gin.H{"error": "request not exist"})
		// 	return
		// }
		c.JSON(http.StatusOK, result)
	}
}

func GetRentListOfUserById(user_id string)([]bson.M,error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "user_id", Value: user_id},
		}},
	}
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "books"},
			{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "book"},
		}},
	}
	lookupStage2 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "let", Value: bson.D{{Key: "user_id", Value: "$user_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$user_id", "$$user_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "user"},
		}},
	}
	lookupStage3 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "book_detail"},
			{Key: "let", Value: bson.D{{Key: "book_detail_id", Value: "$book_detail_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_detail_id", "$$book_detail_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "book_detail"},
		}},
	}
	lookupStage4 := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "book_types"},
			{Key: "let", Value: bson.D{{Key: "book", Value: "$book"}}},
			{Key: "pipeline", Value: bson.A{
				// bson.D{{Key: "$unwind",Value: "$$book"}},
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$typeid", "$$book.type_id"}}}}}}},
				bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
			}},
			{Key: "as", Value: "type"},
		}},
	}

	unwindStage := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
	}
	unwindStage2 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$user"},{Key: "preserveNullAndEmptyArrays",Value: true}}},
	}
	unwindStage3 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$book_detail"}}},
	}
	unwindStage4 := bson.D{
		{Key: "$unwind",Value: bson.D{{Key: "path",Value: "$type"}}},
	}
	// setStage := bson.D{
	// 	{"$set", bson.D{
	// 		{"book", "$book.0"},
	// 		{"user", "$user.0"},
	// 	}},
	// }
	// projectStage := bson.D{
	// 	{"$project", bson.D{
	// 		{"_id", 0},
	// 		{"user",0},
	// 		{"book",0},
	// 	}},
	// }
	cursor, err := BookRentCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, lookupStage2,lookupStage3,unwindStage,lookupStage4,unwindStage2,unwindStage3,unwindStage4,
	})
	result := []bson.M{}
	if err != nil {
		return result,err
	}
	cursor.All(ctx, &result)
	// if len(result) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{"error": "request not exist"})
	// 	return
	// }
	return result,nil
}