package controllers

import (
	"context"
	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/baguette/go-lib/database"
	"github.com/baguette/go-lib/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/matoous/go-nanoid/v2"
)

var BookCollection = database.OpenCollection(database.Client, "books")
var HistoryCollection = database.OpenCollection(database.Client, "histories")
var BookBorrowedCollection = database.OpenCollection(database.Client, "book_borrow")
var BookRentCollection = database.OpenCollection(database.Client, "book_rent")
var BookDetailCollection = database.OpenCollection(database.Client, "book_detail")
var TypeCollections = database.OpenCollection(database.Client, "book_types")

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		books := []models.BookModel{}
		cursor, cursorErr := BookCollection.Find(ctx, bson.D{{
			Key:   "",
			Value: nil,
		}})
		if cursorErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": cursorErr.Error()})
			return
		}
		if err := cursor.All(ctx, &books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}
func GetBooks2() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		page, _ := strconv.Atoi(c.Query("page"))
		// if page <= 0 {
		// 	page = 1
		// }
		recordPerPage := 10
		startIndex := (page - 1) * recordPerPage
		books := []bson.M{}
		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$type"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, skipStage, limitStage, lookupStage, unwindStage, unsetStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := cursor.All(ctx, &books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}
func GetBookByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		id := c.Param("book_id")
		bookModel := models.BookModel{}
		if err := BookCollection.FindOne(ctx, bson.D{{Key: "book_id", Value: id}}).Decode(&bookModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bookModel)
	}
}

// books/search
func GetBooksByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		books := []bson.M{}
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 10
		startIndex := (page - 1) * recordPerPage
		name := c.Query("search_text")
		// query := fmt.Sprintf("%s", name)
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "name", Value: bson.D{{Key: "$regex", Value: name}, {Key: "$options", Value: "i"}}},
			}},
		}
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$type"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, skipStage, limitStage, lookupStage, unwindStage, unsetStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := cursor.All(ctx, &books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		bookModel := models.BookModel{}
		if err := c.ShouldBindJSON(&bookModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad book model"})
			return
		}
		bookModel.Id = primitive.NewObjectID()
		bookModel.Book_id, _ = gonanoid.Generate(NanoidString, 12)
		insertRes, insertErr := BookCollection.InsertOne(ctx, bookModel)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting book"})
			return
		}
		c.JSON(http.StatusOK, insertRes)
	}
}

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		id := c.Param("book_id")
		bookModel := models.BookModel{}
		if err := c.BindJSON(&bookModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad book model"})
			return
		}
		updateObj := bson.D{}
		if bookModel.Book_img != "" {
			updateObj = append(updateObj, bson.E{Key: "book_img", Value: bookModel.Book_img})
		}
		if bookModel.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: bookModel.Name})
		}
		if bookModel.Type_id != "" {
			updateObj = append(updateObj, bson.E{Key: "type_id", Value: bookModel.Type_id})
		}
		updateRes, updateErr := BookCollection.UpdateOne(ctx, bson.M{"book_id": id}, bson.D{{Key: "$set", Value: updateObj}})
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update"})
			return
		}
		c.JSON(http.StatusOK, updateRes)
	}
}

//	func GetLatestBooks() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
//			defer cancel()
//			latestBooks := []bson.M{}
//			page, _ := strconv.Atoi(c.Query("page"))
//			if page <= 0 {
//				page = 1
//			}
//			recordPerPage := 30
//			skip := (page - 1) * recordPerPage
//			opts := options.Find().SetSort(bson.M{"updated_at": -1}).SetLimit(int64(recordPerPage)).SetSkip(int64(skip))
//			cursor, err := BookCollection.Find(ctx, bson.M{}, opts)
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere"})
//				return
//			}
//			if err := cursor.All(ctx, &latestBooks); err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere 2"})
//				return
//			}
//			c.JSON(http.StatusOK, latestBooks)
//		}
//	}
func GetNewestBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		latestBooks := []bson.M{}
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 30
		startIndex := (page - 1) * recordPerPage
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{},
			}},
		}
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$type"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		sortStage := bson.D{
			{Key: "$sort", Value: bson.D{{Key: "updated_at", Value: -1}}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, skipStage, limitStage, sortStage, lookupStage, unwindStage, unsetStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere"})
			return
		}
		if err := cursor.All(ctx, &latestBooks); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere 2"})
			return
		}
		c.JSON(http.StatusOK, latestBooks)
	}
}
func GetPopularBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		popularBooks := []bson.M{}
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 30
		startIndex := (page - 1) * recordPerPage
		// matchStage := bson.D{
		// 	{Key: "$match", Value: bson.D{
		// 		{},
		// 	}},
		// }
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$type"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		sortStage := bson.D{
			{Key: "$sort", Value: bson.D{{Key: "borrowed_quantity", Value: -1}}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			skipStage, limitStage, sortStage, lookupStage, unwindStage, unsetStage,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere"})
			return
		}
		if err := cursor.All(ctx, &popularBooks); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere 2"})
			return
		}
		c.JSON(http.StatusOK, popularBooks)
	}
}
func GetBooksByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		book_type := c.Query("book_type")
		typeModel := models.BookTypeModel{}
		if err := TypeCollections.FindOne(ctx, bson.M{"typename": book_type}).Decode(&typeModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
			return
		}
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 20
		startIndex := (page - 1) * recordPerPage
		books := []bson.M{}
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "type_id", Value: typeModel.TypeId},
			}},
		}
		skipStage := bson.D{{Key: "$skip", Value: startIndex}}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}
		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$type"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		}
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: recordPerPage}}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, skipStage, limitStage, lookupStage, unwindStage, unsetStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error aggregating"})
			return
		}
		cursor.All(ctx, &books)
		c.JSON(http.StatusOK, books)

	}
}

func FindBookToRentWithId(book_id string) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	bookToRent := bson.M{}
	// bookModel := models.BookModel{}
	bookDetail := models.BookDetailModel{}
	// BookCollection.FindOne(ctx, bson.M{"book_id": book_id}).Decode(&bookModel)
	if err := BookDetailCollection.FindOne(ctx, bson.D{
		{Key: "$and",Value: bson.A{
			bson.D{{Key: "book_id",Value: book_id}},
			bson.D{{Key: "status",Value: "ready"}},
		}},
	}).Decode(&bookDetail); err != nil {
		return bookToRent, err
	}
	bookToRent["book_id"] = bookDetail.Book_id
	bookToRent["book_detail_id"] = bookDetail.Book_detail_id
	bookToRent["location"] = bookDetail.Location
	return bookToRent, nil
}
func RentABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		obj := bson.M{}
		// userid
		// bookid
		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book_id := obj["book_id"].(string)
		bookToRent, err := FindBookToRentWithId(book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding book to rent"})
			return
		}
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// nextMonth, _ := time.Parse(time.RFC3339, time.Now().Add(30*24*time.Hour).Format(time.RFC3339))
		// rent request
		bookRentModel := models.BookRentModel{}
		bookRentModel.Id = primitive.NewObjectID()
		bookRentModel.Book_id = book_id
		bookRentModel.Reserve_date = now
		bookRentModel.Book_rent_id, _ = gonanoid.Generate(NanoidString, 12)
		bookRentModel.Book_detail_id = bookToRent["book_detail_id"].(string)
		bookRentModel.User_id = obj["user_id"].(string)
		insertRes, insertErr := BookRentCollection.InsertOne(ctx, bookRentModel)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error renting"})
			return
		}
		updateObj := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "booked"},{Key: "updated_at",Value: now}}}}
		//
		BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookToRent["book_detail_id"].(string)}, updateObj)
		c.JSON(http.StatusOK, gin.H{
			"succeeded":      insertRes,
			"user_id":        obj["user_id"].(string),
			"book_detail_id": bookToRent["book_detail_id"].(string),
			"reserve_date":   now,
		})
	}
}

func GetBookDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		bookDetail := []bson.M{}
		book_id := c.Param("book_id")
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}
		recordPerPage := 5
		startIndex := (page - 1) * recordPerPage
		opt := options.Find().SetSkip(int64(startIndex)).SetLimit(int64(recordPerPage))
		cursor, err := BookDetailCollection.Find(ctx, bson.M{"book_id": book_id}, opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding"})
			return
		}
		cursor.All(ctx, &bookDetail)
		c.JSON(http.StatusOK, bookDetail)
	}
}
