package controllers

import (
	"context"
	"fmt"
	"net/http"

	// "fmt"
	// "net/http"
	// "strconv"
	"time"

	"github.com/baguette/go-lib/helper"
	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/matoous/go-nanoid/v2"
)

func HireABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		//rent_id
		//token
		rec := bson.M{}
		c.ShouldBindJSON(&rec)
		book_rent_id := rec["book_rent_id"].(string)
		token := rec["token"].(string)
		claim, msg := helper.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		if claim.Role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
			return
		}
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		week1, _ := time.Parse(time.RFC3339, time.Now().Add(7*24*time.Hour).Format(time.RFC3339))
		obj := bson.M{}
		obj2 := bson.M{}
		if err := BookRentCollection.FindOne(ctx, bson.M{"book_rent_id": book_rent_id}).Decode(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "what up"})
			return
		}
		book_detail_id := obj["book_detail_id"].(string)
		book_id := obj["book_id"].(string)
		user_id := obj["user_id"].(string)
		if err := BookDetailCollection.FindOne(ctx, bson.M{"book_detail_id": book_detail_id}).Decode(&obj2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "what up"})
			return
		}
		//create borrow model
		bookBorrowModel := models.BookBorrowModel{}
		bookBorrowModel.Id = primitive.NewObjectID()
		bookBorrowModel.Book_id = book_id
		bookBorrowModel.User_id = user_id
		bookBorrowModel.Book_detail_id = book_detail_id
		bookBorrowModel.Book_hire_id, _ = gonanoid.Generate(NanoidString, 12)
		bookBorrowModel.Date_borrowed = now
		bookBorrowModel.Date_end = week1

		//update book detail status
		updateObj := bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: "borrowed"}, {Key: "updated_at", Value: now}}},
		}

		//create history borrowing
		historyModel := models.HistoryModel{}
		historyModel.Id = primitive.NewObjectID()
		historyModel.Book_id = book_id
		historyModel.User_id = user_id
		historyModel.Book_detail_id = book_detail_id
		historyModel.Date_borrowed = now
		historyModel.Date_return = week1
		historyModel.History_id, _ = gonanoid.Generate(NanoidString, 12)
		historyModel.Book_hire_id = bookBorrowModel.Book_hire_id
		historyModel.Status = "borrowing"
		BookRentCollection.DeleteOne(ctx, bson.M{"book_rent_id": book_rent_id})
		BookBorrowedCollection.InsertOne(ctx, bookBorrowModel)
		BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookBorrowModel.Book_detail_id}, updateObj)
		HistoryCollection.InsertOne(ctx, historyModel)
		c.JSON(http.StatusOK, gin.H{
			"status":         "success",
			"history_id":     historyModel.History_id,
			"book_detail_id": bookBorrowModel.Book_detail_id,
			"location":       obj2["location"].(string),
			"book_id":        book_id,
			"date_borrowed":  now.Local(),
			"date_end":       week1.Local(),
			"book_hire_id":   bookBorrowModel.Book_hire_id,
		})
	}
}

func GetBorrowList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		borrowList := []bson.M{}
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
		cursor, _ := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
		})
		cursor.All(ctx, &borrowList)
		c.JSON(http.StatusOK, borrowList)
	}
}

func GetBorrowListOfABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		borrowList := []bson.M{}
		book_id := c.Param("book_id")
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "book_id", Value: book_id}}}}
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
		// projectStage := bson.D{
		// 	{Key: "$project",Value: bson.D{
		// 		{Key: "_id",Value: 0},
		// 		{Key: "user._id",Value: 0},
		// 		{Key: "book._id",Value: 0},
		// 		{Key: "user.password",Value: 0},
		// 		{Key: "user.email",Value: 0},
		// 		{Key: "user.phone",Value: 0},
		// 		{Key: "user.role",Value: 0},
		// 		{Key: "user.created_at",Value: 0},
		// 		{Key: "user.updated_at",Value: 0},
		// 		{Key:"user_id",Value:0},
		// 	}},
		// }
		unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$book"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
		unwindStage2 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				// {Key: "book.book_id",Value: 1},
				// {Key: "book.book_img",Value: 1},
				// {Key: "book.name",Value: 1},
				// {Key: "book.book_detail_id",Value: 1},
				// {Key: "book.book_hire_id",Value: 1},
				// {Key: "book_id",Value: 1},
				// {Key: "book_id",Value: 1},
				// {Key: "date_borrowed",Value: 1},
				// {Key: "date_end",Value: 1},
				// {Key: "user.user_id",Value: 1},
				// {Key: "user.name",Value: 1},
			}},
		}
		cursor, _ := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, lookupStage2, unwindStage, unwindStage2, projectStage,
		})
		cursor.All(ctx, &borrowList)
		c.JSON(http.StatusOK, borrowList)
	}
}

func ReturnABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		rec := bson.M{}
		c.ShouldBindJSON(&rec)
		token := rec["token"].(string)
		claim, msg := helper.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		if claim.Role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
			return
		}
		book_hire_id := c.Param("book_hire_id")
		bookBorrowModel := bson.M{}
		BookBorrowedCollection.FindOneAndDelete(ctx, bson.M{"book_hire_id": book_hire_id}).Decode(&bookBorrowModel)
		updateObj := bson.M{"$set": bson.M{"status": "ready", "updated_at": now}}
		updateHis := bson.M{"$set": bson.M{"status": "returned", "date_return": now, "book_hire_id": ""}}
		updateObj2 := bson.M{"$inc": bson.M{"amount": 1, "borrowed_quantity": 1}}
		_, err1 := BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookBorrowModel["book_detail_id"].(string)}, updateObj)
		_, err2 := HistoryCollection.UpdateOne(ctx, bson.M{"book_hire_id": bookBorrowModel["book_hire_id"].(string)}, updateHis)
		_, err3 := BookCollection.UpdateOne(ctx, bson.M{"book_id": bookBorrowModel["book_id"].(string)}, updateObj2)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error 1"})
			return
		}
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error 2"})
			return
		}
		if err3 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error 3"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "returned",
		})
	}
}

func BookBorrowById(id string) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// history := bson.M{}
	// HistoryCollection.FindOne(ctx, bson.M{"history_id": history_id}).Decode(&history)
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "book_hire_id", Value: id},
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
	cursor, err := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, lookupStage2, lookupStage3, unwindStage, lookupStage4, unwindStage2, unwindStage3, unwindStage4,
	})
	res := []bson.M{}
	if err != nil {
		return bson.M{}, err
	}
	cursor.All(ctx, &res)
	return res[0], nil
}

func GetBookBorrowById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		book_hire_id := c.Query("search_id")
		// history := bson.M{}
		// HistoryCollection.FindOne(ctx, bson.M{"history_id": history_id}).Decode(&history)
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "book_hire_id", Value: book_hire_id},
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
		cursor, err := BookBorrowedCollection.Aggregate(ctx, mongo.Pipeline{
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

func DirectlyBorrow(email string, book_id string) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	user := models.UserModel{}
	result := bson.M{}
	if err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		newErr := fmt.Errorf("user not found")
		return result, newErr
	}
	count, _ := BookRentCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id})
	if count > 0 {
		err := fmt.Errorf("already borrowing another books")
		return result, err
	}
	count2, _ := BookBorrowedCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id})
	if count2 > 0 {
		err := fmt.Errorf("already borrowing another books")
		return result, err
	}
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	week1, _ := time.Parse(time.RFC3339, time.Now().Add(7*24*time.Hour).Format(time.RFC3339))
	bookBorrowModel := models.BookBorrowModel{}
	book, bookErr := FindBookToRentWithId(book_id)
	if bookErr != nil {
		return result, bookErr
	}
	/// book["location"].(string)
	bookBorrowModel.Book_id = book_id
	bookBorrowModel.Book_detail_id = book["book_detail_id"].(string)
	bookBorrowModel.Date_borrowed = now
	bookBorrowModel.Date_end = week1
	bookBorrowModel.Id = primitive.NewObjectID()
	bookBorrowModel.User_id = user.User_id
	bookBorrowModel.Book_hire_id, _ = gonanoid.Generate(NanoidString, 12)

	///
	updateObj := bson.D{
		{Key: "$set", Value: bson.D{{Key: "status", Value: "borrowed"}, {Key: "updated_at", Value: now}}},
	}
	updateObj2 := bson.M{"$inc": bson.M{"amount": -1}}
	///
	historyModel := models.HistoryModel{}
	historyModel.Id = primitive.NewObjectID()
	historyModel.History_id, _ = gonanoid.Generate(NanoidString, 12)
	historyModel.Date_borrowed = now
	historyModel.Date_return = week1
	historyModel.Book_id = book_id
	historyModel.User_id = user.User_id
	historyModel.Book_detail_id = book["book_detail_id"].(string)
	historyModel.Status = "borrowing"
	historyModel.Book_hire_id = bookBorrowModel.Book_hire_id
	///
	BookBorrowedCollection.InsertOne(ctx, bookBorrowModel)
	BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookBorrowModel.Book_detail_id}, updateObj)
	BookCollection.UpdateOne(ctx, bson.M{"book_id": book_id}, updateObj2)
	HistoryCollection.InsertOne(ctx, historyModel)
	// c.JSON(http.StatusOK, gin.H{
	// 		"status":         "success",
	// 		"history_id":     historyModel.History_id,
	// 		"book_detail_id": bookBorrowModel.Book_detail_id,
	// 		"book_id":        bookBorrowModel.Book_id,
	// 		"date_borrowed":  bookBorrowModel.Date_borrowed,
	// 		"book_hire_id":   bookBorrowModel.Book_hire_id,
	// 	})
	// result["book_hire_id"] = bookBorrowModel.Book_hire_id
	// result["status"] = "success"
	// result[""]
	result = bson.M{
		"status":        "success",
		"book_id":       book_id,
		"book_hire_id":  bookBorrowModel.Book_hire_id,
		"history_id":    historyModel.History_id,
		"date_borrowed": bookBorrowModel.Date_borrowed,
		"date_end":      bookBorrowModel.Date_end,
		"user_id":       user.User_id,
	}
	return result, nil
}

func DirectlyBorrowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rec := bson.M{}
		c.BindJSON(&rec)
		email := rec["email"].(string)
		book_id := rec["book_id"].(string)
		result, err := DirectlyBorrow(email, book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
