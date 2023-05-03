package seeds

import (
	"context"
	"time"

	"github.com/baguette/seed-lib/database"
	"github.com/baguette/seed-lib/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var UserCollection = database.OpenCollection(database.Client, "users")
var BookCollection = database.OpenCollection(database.Client, "books")
var HistoryCollection = database.OpenCollection(database.Client, "histories")
var BookBorrowCollection = database.OpenCollection(database.Client, "book_borrow")
var BookRentCollection = database.OpenCollection(database.Client, "book_rent")
var BookDetailCollection = database.OpenCollection(database.Client, "book_detail")
var TypeCollections = database.OpenCollection(database.Client, "book_types")
var NanoidString = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func SeedUser() {
	// ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	// defer cancel()
	User1 := models.UserModel{
		Name:     "user1",
		Email:    "user1@gmail.com",
		Phone:    "1",
		Password: "$2a$10$BlbYzo6uJGyLLowPnP3sKuZc6xO.vk2LxDvYmvt73/N5ummZrxiJS",
		Role:     "admin",
	}
	User2 := models.UserModel{
		Name:     "user2",
		Email:    "user2@gmail.com",
		Phone:    "2",
		Password: "$2a$10$BlbYzo6uJGyLLowPnP3sKuZc6xO.vk2LxDvYmvt73/N5ummZrxiJS",
		Role:     "user",
	}
	User3 := models.UserModel{
		Name:     "user3",
		Email:    "user3@gmail.com",
		Phone:    "3",
		Password: "$2a$10$BlbYzo6uJGyLLowPnP3sKuZc6xO.vk2LxDvYmvt73/N5ummZrxiJS",
		Role:     "user",
	}
	User4 := models.UserModel{
		Name:     "user4",
		Email:    "user4@gmail.com",
		Phone:    "4",
		Password: "$2a$10$BlbYzo6uJGyLLowPnP3sKuZc6xO.vk2LxDvYmvt73/N5ummZrxiJS",
		Role:     "user",
	}
	User5 := models.UserModel{
		Name:     "user5",
		Email:    "user5@gmail.com",
		Phone:    "5",
		Password: "$2a$10$BlbYzo6uJGyLLowPnP3sKuZc6xO.vk2LxDvYmvt73/N5ummZrxiJS",
		Role:     "user",
	}
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	User1.User_id = "djWzBNgD7Cvc"
	User2.User_id = "6ZrgtsjaKQi7"
	User3.User_id = "6clomWjyrM1n"
	User4.User_id = "UtkvLPlXNzm4"
	User5.User_id = "ZzcPF1BP8t2F"
	
	User1.Id= primitive.NewObjectID()
	User2.Id = primitive.NewObjectID()
	User3.Id = primitive.NewObjectID()
	User4.Id = primitive.NewObjectID()
	User5.Id = primitive.NewObjectID()
	User1.Created_at = now
	User2.Created_at = now
	User3.Created_at = now
	User4.Created_at = now
	User5.Created_at = now
	User1.Updated_at = now
	User2.Updated_at = now
	User3.Updated_at = now
	User4.Updated_at = now
	User5.Updated_at = now
	arr := []interface{}{
		User1,
		User2,
		User3,
		User4,
		User5,
	}	
	UserCollection.InsertMany(context.TODO(),arr)
}
