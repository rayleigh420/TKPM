package controllers

import (
	"context"
	"math"
	"strings"

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
		books := []bson.M{}
		// lookupStage := bson.D{
		// 	{Key: "$lookup", Value: bson.D{
		// 		{Key: "from", Value: "book_detail"},
		// 		{Key: "let", Value: bson.D{{Key: "book_id", Value: "$book_id"}}},
		// 		{Key: "pipeline", Value: bson.A{
		// 			bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$book_id", "$$book_id"}}}}}}},
		// 			bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}}}},
		// 		}},
		// 		{Key: "as", Value: "book_detail"},
		// 	}},
		// }
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_detail"},
				{Key: "localField", Value: "book_id"},
				{Key: "foreignField", Value: "book_id"},
				{Key: "as", Value: "book_detail"},
			}},
		}
		lookupStage2 := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "book_types"},
				{Key: "localField", Value: "type_id"},
				{Key: "foreignField", Value: "typeid"},
				{Key: "as", Value: "type"},
			}},
		}

		unwindStage := bson.D{
			{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$type"}}},
		}
		cursor, cursorErr := BookCollection.Aggregate(ctx, mongo.Pipeline{
			lookupStage, lookupStage2, unwindStage,
		})
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

		// group := bson.D{
		// 	{Key: "$group",Value: bson.D{
		// 		{Key: "_id",Value: "null"},
		// 		{Key: "count",Value: bson.D{{Key: "$sum",Value: 1}}},
		// 	}},
		// }
		// project := bson.D{
		// 	{Key: "$project",Value: bson.D{{Key: "_id",Value: 0}}},
		// }
		// cursor,_ := BookCollection.Aggregate(ctx,mongo.Pipeline{
		// 	group,project,
		// })

		count, _ := BookCollection.CountDocuments(ctx, bson.D{{}})
		count2 := float64(count)
		count2 = math.Ceil(count2 / 10)
		totalCount := strconv.Itoa(int(count2))
		c.Writer.Header().Add("Total", totalCount)
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
		bookModel := bson.M{}
		if err := BookCollection.FindOne(ctx, bson.D{{Key: "book_id", Value: id}}).Decode(&bookModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		typeModel := bson.M{}
		TypeCollections.FindOne(ctx, bson.M{"typeid": bookModel["type_id"].(string)}).Decode(&typeModel)
		booksRelated := []bson.M{}
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "book_id", Value: bson.D{{Key: "$ne", Value: id}}},
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "type_id", Value: bookModel["type_id"].(string)}},
					bson.D{{Key: "author", Value: bson.D{{Key: "$regex", Value: bookModel["author"].(string)}, {Key: "$options", Value: "i"}}}},
				}},
			}},
		}
		// matchStage2 := bson.D{
		// 	{"$match",bson.D{
		// 		{"book_id",bson.D{{"$not",id}}},
		// 	}},
		// }
		// matchStage := bson.D{
		// 	{Key: "$match", Value: bson.D{
		// 		{Key: "$and", Value: bson.A{
		// 			bson.D{{Key: "book_id", Value: bson.D{{Key: "$not", Value: id}}}},
		// 			bson.D{{Key: "$or", Value: bson.A{
		// 				bson.D{{Key: "type_id", Value: bookModel["type_id"].(string)}},
		// 				bson.D{{Key: "author", Value: bson.D{{Key: "$regex", Value: bookModel["author"].(string)}, {Key: "$options", Value: "i"}}}},
		// 			}}},
		// 		}},
		// 	}},
		// }
		unsetStage := bson.D{{Key: "$unset", Value: bson.A{
			"type._id", "type_id", "_id", "created_at", "description", "details", "license", "page",
		}}}
		limitStage := bson.D{{Key: "$limit", Value: 10}}
		cursor, _ := BookCollection.Aggregate(ctx, mongo.Pipeline{
			limitStage, matchStage, unsetStage,
		})
		cursor.All(ctx, &booksRelated)
		bookModel["type_name"] = typeModel["typename"].(string)
		if len(booksRelated) != 0 {
			bookModel["related_books"] = booksRelated
		} else {
			bookModel["related_books"] = []bson.M{}
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
		rec := bson.M{}
		c.Bind(&rec)

		name := rec["search_text"].(string)
		// query := fmt.Sprintf("%s", name)
		count, _ := BookCollection.CountDocuments(ctx, bson.D{
			{Key: "name", Value: bson.D{{Key: "$regex", Value: name}, {Key: "$options", Value: "i"}}},
		})
		count2 := float64(count)
		count2 = math.Ceil(count2 / 10)
		totalCount := strconv.Itoa(int(count2))
		c.Writer.Header().Add("Total", totalCount)
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
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// bookModel := models.BookModel{}
		rec := bson.M{}
		if err := c.ShouldBindJSON(&rec); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad book model"})
			return
		}
		opt := options.FindOne().SetSort(bson.M{"$natural": -1})
		res := models.BookModel{}
		BookCollection.FindOne(ctx, bson.M{}, opt).Decode(&res)
		co := strings.Split(res.Book_id, "B")[1]
		count, _ := strconv.Atoi(co)
		count++
		counts := strconv.Itoa(int(count))
		id := "B" + counts
		// bookModel.Id = primitive.NewObjectID()
		// bookModel.Book_id = id
		// bookModel.Created_at = now
		// bookModel.Updated_at = now
		typeid, err := InsertTypeByName(rec["type_name"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error type"})
			return
		}
		delete(rec, "type_name")
		rec["_id"] = primitive.NewObjectID()
		rec["book_id"] = id
		rec["created_at"] = now
		rec["updated_at"] = now
		rec["type_id"] = typeid

		// bookModel.Name = rec["name"].(string)
		// bookModel.Publisher = rec["publisher"].(string)
		// bookModel.Yearpublished = rec["yearpublished"].(int64)
		// bookModel.
		// bookModel.Type_id = typeid
		// bookModel.Amount = rec["amount"].(int64)
		// bookModel.Author = rec["author"].(string)
		// bookModel.Book_img = rec["book_image"].(string)
		// bookModel.Borrowed_quantity = rec["borrowed_quantity"].(int64)
		// bookModel.Description = rec["description"].(string)
		// bookModel.Details = rec["details"].(string)
		// bookModel.License = rec["license"].(string)
		// bookModel.Page = rec["page"].(int64)
		// bookModel.Publishing_location = rec["publishing_location"].(string)

		_, insertErr := BookCollection.InsertOne(ctx, rec)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"book_id":   id,
			"book_name": rec["name"].(string),
		})
	}
}

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		id := c.Param("book_id")
		bookModel := bson.M{}
		if err := c.BindJSON(&bookModel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad book model"})
			return
		}
		updateObj := bson.D{}
		if name := bookModel["name"].(string); name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: name})
		}
		if type_name := bookModel["type_name"].(string); type_name != "" {
			typeid, err := InsertTypeByName(type_name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error type"})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "type_id", Value: typeid})
		}
		if des := bookModel["description"].(string); des != "" {
			updateObj = append(updateObj, bson.E{Key: "description", Value: des})
		}
		if det := bookModel["details"].(string); det != "" {
			updateObj = append(updateObj, bson.E{Key: "details", Value: det})
		}
		if lic := bookModel["license"].(string); lic != "" {
			updateObj = append(updateObj, bson.E{Key: "license", Value: lic})
		}
		if a := bookModel["book_image"].(string); a != "" {
			updateObj = append(updateObj, bson.E{Key: "book_image", Value: a})
		}
		if des := bookModel["publisher"].(string); des != "" {
			updateObj = append(updateObj, bson.E{Key: "publisher", Value: des})
		}
		if des := bookModel["author"].(string); des != "" {
			updateObj = append(updateObj, bson.E{Key: "author", Value: des})
		}
		if des := bookModel["yearpublished"].(float64); des != 0 {
			updateObj = append(updateObj, bson.E{Key: "yearpublished", Value: int64(des)})
		}
		if des := bookModel["page"].(float64); des != 0 {
			updateObj = append(updateObj, bson.E{Key: "page", Value: int64(des)})
		}
		if des := bookModel["publishing_location"].(string); des != "" {
			updateObj = append(updateObj, bson.E{Key: "publishing_location", Value: des})
		}

		updateRes, updateErr := BookCollection.UpdateOne(ctx, bson.M{"book_id": id}, bson.D{{Key: "$set", Value: updateObj}})
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update"})
			return
		}
		c.JSON(http.StatusOK, updateRes)
	}
}

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

		count, _ := BookCollection.CountDocuments(ctx, bson.D{{}})
		count2 := float64(count)
		count2 = math.Ceil(count2 / 10)
		totalCount := strconv.Itoa(int(count2))
		c.Writer.Header().Add("Total", totalCount)

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
		count, _ := BookCollection.CountDocuments(ctx, bson.D{{}})
		count2 := float64(count)
		count2 = math.Ceil(count2 / 10)
		totalCount := strconv.Itoa(int(count2))
		c.Writer.Header().Add("Total", totalCount)

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
		count, _ := BookCollection.CountDocuments(ctx, bson.D{{Key: "type_id", Value: typeModel.TypeId}})
		count2 := float64(count)
		count2 = math.Ceil(count2 / 10)
		totalCount := strconv.Itoa(int(count2))
		c.Writer.Header().Add("Total", totalCount)

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
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "book_id", Value: book_id}},
			bson.D{{Key: "status", Value: "ready"}},
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad rent model"})
			return
		}
		//check if user is already borrowing
		count, _ := BookRentCollection.CountDocuments(ctx, bson.M{"user_id": obj["user_id"].(string)})
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
				"error":  "already requesting another book",
			})
			return
		}
		count2, _ := BookBorrowedCollection.CountDocuments(ctx, bson.M{"user_id": obj["user_id"].(string)})
		if count2 > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
				"error":  "already borrowing another book",
			})
			return
		}
		book_id := obj["book_id"].(string)
		bookToRent, err := FindBookToRentWithId(book_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding book to rent"})
			return
		}
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		amount := bson.M{}
		opts := options.FindOne().SetProjection(bson.M{"amount": 1})
		BookCollection.FindOne(ctx, bson.M{"book_id": book_id}, opts).Decode(&amount)
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
		updateObj := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "booked"}, {Key: "updated_at", Value: now}}}}
		// updateObj2 := bson.D{{Key: "$set", Value: bson.D{{Key: "amount", Value: amount["amount"].(int64) - 1}}}}
		upd := bson.M{"$inc":bson.M{"amount":-1}}
		//
		BookCollection.UpdateOne(ctx, bson.M{"book_id": book_id}, upd)
		BookDetailCollection.UpdateOne(ctx, bson.M{"book_detail_id": bookToRent["book_detail_id"].(string)}, updateObj)
		c.JSON(http.StatusOK, gin.H{
			"succeeded":      insertRes,
			"user_id":        obj["user_id"].(string),
			"book_detail_id": bookToRent["book_detail_id"].(string),
			"book_rent_id":   bookRentModel.Book_rent_id,
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
		// page, _ := strconv.Atoi(c.Query("page"))
		// if page <= 0 {
		// 	page = 1
		// }
		// recordPerPage := 5
		// startIndex := (page - 1) * recordPerPage
		// count, _ := BookDetailCollection.CountDocuments(ctx, bson.D{{Key: "book_id", Value: book_id}})
		// count2 := float64(count)
		// count2 = math.Ceil(count2 / 10)
		// totalCount := strconv.Itoa(int(count2))
		// c.Writer.Header().Add("Total", totalCount)

		// opt := options.Find().SetSkip(int64(startIndex)).SetLimit(int64(recordPerPage))
		cursor, err := BookDetailCollection.Find(ctx, bson.M{"book_id": book_id})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding"})
			return
		}
		cursor.All(ctx, &bookDetail)
		c.JSON(http.StatusOK, bookDetail)
	}
}

func DeleteBook(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	rec := bson.M{}
	err := BookCollection.FindOneAndDelete(ctx, bson.M{"book_id": id}).Decode(&rec)
	if err != nil {
		return err
	}
	_, err2 := BookDetailCollection.DeleteMany(ctx, bson.M{"book_id": id})
	BookRentCollection.DeleteMany(ctx, bson.M{"book_id": id})
	BookBorrowedCollection.DeleteMany(ctx, bson.M{"book_id": id})
	HistoryCollection.DeleteMany(ctx, bson.M{"book_id": id})
	if err2 != nil {
		return err
	}
	return nil
}

func DeleteBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("book_id")
		err := DeleteBook(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
