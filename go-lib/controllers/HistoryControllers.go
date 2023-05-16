package controllers

import (
	"context"
	"net/http"
	// "strconv"
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
		// // unsetStage := bson.D{{"$unset",bson.A{"_id"}}}
		// unwindStage := bson.D{
		// 	{Key: "$unwind", Value: bson.D{
		// 		{Key: "path", Value: "$book"},
		// 		{Key: "preserveNullAndEmptyArrays", Value: false},
		// 	}},
		// }
		// unwindStage2 := bson.D{
		// 	{Key: "$unwind", Value: bson.D{
		// 		{Key: "path", Value: "$user"},
		// 		{Key: "preserveNullAndEmptyArrays", Value: false},
		// 	}},
		// }
		// projectStage := bson.D{
		// 	{Key: "$project", Value: bson.D{
		// 		{Key: "book.book_id", Value: 1},
		// 		{Key: "book.book_img", Value: 1},
		// 		{Key: "user.name", Value: 1},
		// 		{Key: "user.avatar", Value: 1},
		// 		{Key: "date_borrowed", Value: 1},
		// 		{Key: "date_return", Value: 1},
		// 		{Key: "user_id", Value: 1},
		// 		{Key: "status", Value: 1},
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
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}},
		}
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage, lookupStage2, lookupStage3, unwindStage, unwindStage2, unwindStage3,
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

func GetHistoryOfUser(user_id string)([]bson.M,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		res := []bson.M{}
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}
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
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$type"}}},
		}
		// projectStage := bson.D{
		// 	{Key: "$project", Value: bson.D{
		// 		{Key: "history_id", Value: 1},
		// 		{Key: "user.avatar", Value: 1},
		// 		{Key: "user.name", Value: 1},
		// 		{Key: "book.book_id", Value: 1},
		// 		{Key: "book.name", Value: 1},
		// 		{Key: "book.book_img", Value: 1},
		// 		{Key: "date_borrowed", Value: 1},
		// 		{Key: "date_return", Value: 1},
		// 	}},
		// }
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
		})
		if err != nil {
			return res,err
		}
		cursor.All(ctx, &res)
		return res,nil
}

func GetHistoryByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		res,err := GetHistoryOfUser(user_id)
		res2,err2 := GetRentListOfUserById(user_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		if err2 != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		res = append(res, res2[0:]...)
		c.JSON(http.StatusOK,res)
	}
}

func GetReturnedBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		res := []bson.M{}
		// cursor,_ := HistoryCollection.Find(ctx,bson.M{"status":"borrowing"})
		// cursor.All(ctx,&res)
		// matchStage := bson.D{
		// 	{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"status", "returned"}}}}}},
		// }
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "status", Value: "returned"},
			}},
		}
		// sortStage := bson.D{
		// 	{Key: "$sort", Value: bson.D{
		// 		{Key: "date_borrowed", Value: -1},
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
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$type"}}},
		}
		// projectStage := bson.D{
		// 	{Key: "$project", Value: bson.D{
		// 		{Key: "history_id", Value: 1},
		// 		{Key: "user.avatar", Value: 1},
		// 		{Key: "user.name", Value: 1},
		// 		{Key: "book.book_id", Value: 1},
		// 		{Key: "book.name", Value: 1},
		// 		{Key: "book.book_img", Value: 1},
		// 		{Key: "date_borrowed", Value: 1},
		// 		{Key: "date_return", Value: 1},
		// 	}},
		// }
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		cursor.All(ctx, &res)
		c.JSON(http.StatusOK, res)
	}
}

func GetHistoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		history_id := c.Query("search_id")
		// history := bson.M{}
		// HistoryCollection.FindOne(ctx, bson.M{"history_id": history_id}).Decode(&history)
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "history_id", Value: history_id},
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
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: true}}},
		}
		unwindStage2 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: true}}},
		}
		unwindStage3 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}},
		}
		unwindStage4 := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$type"}}},
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
		cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		result := []bson.M{}
		cursor.All(ctx, &result)
		if len(result) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "not exist"})
			return
		}
		c.JSON(http.StatusOK, result[0])
	}
}

func GetHistoryOfBook(book_id string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "book_id", Value: book_id},
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
		{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: true}}},
	}
	unwindStage2 := bson.D{
		{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: true}}},
	}
	unwindStage3 := bson.D{
		{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book_detail"}}},
	}
	unwindStage4 := bson.D{
		{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$type"}}},
	}
	cursor, err := HistoryCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
	})
	if err != nil {
		return []bson.M{}, err
	}
	result := []bson.M{}
	cursor.All(ctx, &result)
	return result, nil
}

func GetHistoryByBookId() gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id := c.Param("book_id")
		res, err := GetHistoryOfBook(book_id)
		res2, err2 := GetRentListOfBook2(book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		res = append(res, res2[0:]...)
		c.JSON(http.StatusOK, res)
	}
}
