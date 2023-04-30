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
		cursor, cursorErr := BookCollection.Find(ctx, bson.D{{}})
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
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "book_types"},
				{"localField", "type_id"},
				{"foreignField", "typeid"},
				{"as", "type"},
			}},
		}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "id", Value: "null"}}},
				{Key: "items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "bookItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$items", startIndex, recordPerPage}}}},
			}},
		}
		projectStage2 := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "bookItems.type._id", Value: 0},
				{Key: "bookItems.type_id", Value: 0},
			}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, lookupStage, groupStage, projectStage, projectStage2,
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
		id := c.Param("id")
		bookModel := models.BookModel{}
		if err := BookCollection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&bookModel); err != nil {
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
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "book_types"},
				{"localField", "type_id"},
				{"foreignField", "typeid"},
				{"as", "type"},
			}},
		}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "bookItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$items", startIndex, recordPerPage}}}},
			}},
		}
		projectStage2 := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "bookItems.type._id", Value: 0},
				{Key: "bookItems.type_id", Value: 0},
			}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,lookupStage, groupStage, projectStage,projectStage2,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err = cursor.All(ctx, &books); err != nil {
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
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		updateRes, updateErr := BookCollection.UpdateOne(ctx, bson.M{"book_id": id}, updateObj, &opt)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update"})
			return
		}
		c.JSON(http.StatusOK, updateRes)
	}
}

// func GetLatestBooks() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
// 		defer cancel()
// 		latestBooks := []bson.M{}
// 		page, _ := strconv.Atoi(c.Query("page"))
// 		if page <= 0 {
// 			page = 1
// 		}
// 		recordPerPage := 30
// 		skip := (page - 1) * recordPerPage
// 		opts := options.Find().SetSort(bson.M{"updated_at": -1}).SetLimit(int64(recordPerPage)).SetSkip(int64(skip))
// 		cursor, err := BookCollection.Find(ctx, bson.M{}, opts)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere"})
// 			return
// 		}
// 		if err := cursor.All(ctx, &latestBooks); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error somewhere 2"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, latestBooks)
// 	}
// }
func GetLatestBooks() gin.HandlerFunc {
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
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "book_types"},
				{"localField", "type_id"},
				{"foreignField", "typeid"},
				{"as", "type"},
			}},
		}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "bookItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$items", startIndex, recordPerPage}}}},
			}},
		}
		projectStage2 := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "bookItems.type._id", Value: 0},
				{Key: "bookItems.type_id", Value: 0},
			}},
		}
		sortStage := bson.D{
				{Key:"$sort",Value:bson.D{{Key:"updated_at",Value:-1}}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,sortStage,lookupStage, groupStage, projectStage,projectStage2,
			
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
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{},
			}},
		}
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "book_types"},
				{"localField", "type_id"},
				{"foreignField", "typeid"},
				{"as", "type"},
			}},
		}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "bookItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$items", startIndex, recordPerPage}}}},
			}},
		}
		projectStage2 := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "bookItems.type._id", Value: 0},
				{Key: "bookItems.type_id", Value: 0},
			}},
		}
		sortStage := bson.D{
				{Key:"$sort",Value:bson.D{{Key:"borrowed_quantity",Value:-1}}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,sortStage,lookupStage, groupStage, projectStage,projectStage2,
			
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
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "book_types"},
				{"localField", "type_id"},
				{"foreignField", "typeid"},
				{"as", "type"},
			}},
		}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "bookItems", Value: bson.D{{Key: "$slice", Value: []interface{}{"$items", startIndex, recordPerPage}}}},
			}},
		}
		projectStage2 := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "bookItems.type._id", Value: 0},
				{Key: "bookItems.type_id", Value: 0},
			}},
		}
		cursor, err := BookCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,lookupStage, groupStage, projectStage,projectStage2,
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
	bookModel := models.BookModel{}
	bookDetail := models.BookDetailModel{}
	BookCollection.FindOne(ctx, bson.M{"book_id": book_id}).Decode(&bookModel)
	if err := BookDetailCollection.FindOne(ctx, bson.M{"book_id": book_id, "status": "ready"}).Decode(&bookDetail); err != nil {
		return bookToRent, err
	}
	bookToRent["book_id"] = bookModel.Book_id
	bookToRent["book_detail_id"] = bookDetail.Book_detail_id
	bookToRent["location"] = bookDetail.Location
	return bookToRent, nil
}
func RentABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		obj := map[string]interface{}{}
		// userid
		// bookid
		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// book_id := obj["book_id"].(string)
		book_id := c.Param("book_id")

		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// nextMonth, _ := time.Parse(time.RFC3339, time.Now().Add(30*24*time.Hour).Format(time.RFC3339))
		// rent request
		bookRentModel := models.BookRentModel{}
		bookRentModel.Id = primitive.NewObjectID()
		bookRentModel.Book_id = book_id
		bookRentModel.Reserve_date = now
		bookRentModel.Book_rent_id, _ = gonanoid.Generate(NanoidString, 12)
		bookRentModel.User_id = obj["user_id"].(string)
		insertRes, insertErr := BookRentCollection.InsertOne(ctx, bookRentModel)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error renting"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"succeeded": insertRes})
	}
}

// func GetLocationOfBook(book_id string) (string, error) {
// 	objectBook, err := FindBookWithId(book_id)
// 	if err != nil {
// 		return "", err
// 	}
// 	return objectBook["book-detail-items"].(map[string]interface{})["location"].(string), nil
// }
