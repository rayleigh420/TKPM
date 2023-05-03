package seeds

import (
	"context"
	"time"

	"github.com/baguette/seed-lib/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var now, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
var week1, _ = time.Parse(time.RFC3339, time.Now().Add(-7*24*time.Hour).Format(time.RFC3339))

func SeedBook() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	Book1 := models.BookModel{
		Name:                "Không Diệt Không Sinh Đừng Sợ Hãi",
		Publisher:           "Saigon Books",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/352/200/300.jpg?hmac=JRE6d4eB1tvPUpBESG8XEM2_22EaXNe2luRrVkydr2E",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   12,
		Amount:              50,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book2 := models.BookModel{
		Name:                "Muôn Kiếp Nhân Sinh - Many Times, Many Lives",
		Publisher:           "NXB Tổng Hợp TPHCM",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/94/200/300.jpg?hmac=CA7x5S3EdSeRqM9TK0RxiKTcx1R96lIncvKTjTP3beI",
		Author:              "Nguyên Phong",
		Borrowed_quantity:   120,
		Amount:              50,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Nguyên Phong",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book3 := models.BookModel{
		Name:                "Muốn An Được An",
		Publisher:           "Thái Hà",
		Yearpublished:       2021,
		Book_img:            "https://fastly.picsum.photos/id/882/200/300.jpg?hmac=j02hkJZfREXyROkSfM6KpIIPHsAaGLmWylVEr5sRzok",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   20,
		Amount:              50,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book4 := models.BookModel{
		Name:                "Tĩnh Lặng - Sức Mạnh Tĩnh Lặng Trong Thế Giới Huyền Ảo",
		Publisher:           "Thái Hà",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/25/200/300.jpg?hmac=ScdLbPfGd_kI3MUHvJUb12Fsg1meDQEaHY_mM613BVM",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   42,
		Amount:              50,
		Type_id:             "2",
		Publishing_location: "Korea",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book5 := models.BookModel{
		Name:                "Chainsaw Man - Tập 11",
		Publisher:           "NXB Trẻ",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/200/200/300.jpg?hmac=XVCLpc2Ddr652IrKMt3L7jISDD4au5O9ZIr3fwBtxo8",
		Author:              "Tatsuki Fujimoto",
		Borrowed_quantity:   512,
		Amount:              50,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                100,
		License:             "Tatsuki Fujimoto",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book6 := models.BookModel{
		Name:                "Khu Rừng Đom Đóm",
		Publisher:           "NXB Trẻ",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/461/200/300.jpg?hmac=dIwmQxeVflRD0QrOZ_p0_q-LpAY7sVhua6FCEIR_xi8",
		Author:              "Yuki Midorikawa",
		Borrowed_quantity:   162,
		Amount:              50,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                300,
		License:             "Yuki Midorikawa",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book7 := models.BookModel{
		Name:                "Cô Gái Nơi Xứ Ngoài",
		Publisher:           "IPM",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/615/200/300.jpg?hmac=ehJCfeXO1-ZbwBXgbYKroA97kTtoPKNoyEbCxnzsYfU",
		Author:              "Nagabe",
		Borrowed_quantity:   127,
		Amount:              50,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                352,
		License:             "Nagabe",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book8 := models.BookModel{
		Name:                "Cô Gái Nơi Xứ Ngoài - Tập 2",
		Publisher:           "IPM",
		Yearpublished:       2001,
		Book_img:            "https://fastly.picsum.photos/id/183/200/300.jpg?hmac=Z9yCtuuIPn5CuOhwIntNEQFIRotghuBn06nqOSL828c",
		Author:              "Nagabe",
		Borrowed_quantity:   182,
		Amount:              50,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                451,
		License:             "Nagabe",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book9 := models.BookModel{
		Name:                "Blockchain - Bản Chất Của Blockchain",
		Publisher:           "1980 Books",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/435/200/300.jpg?hmac=gSfIK7r5rB4f3aJq0RTylep967d4sBRbIUuAOuq433o",
		Author:              "Mark Gates",
		Borrowed_quantity:   192,
		Amount:              50,
		Type_id:             "3",
		Publishing_location: "USA",
		Page:                256,
		License:             "Mark Gates",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book10 := models.BookModel{
		Name:                "Object Oriented Programming",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/972/200/300.jpg?hmac=UMf5f6BV9GkLiz0Xz9kMwm1riiTtlpIG2jt0WrxZ51Q",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   1012,
		Amount:              50,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                444,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book1.Book_id = "B1"
	Book2.Book_id = "B2"
	Book3.Book_id = "B3"
	Book4.Book_id = "B4"
	Book5.Book_id = "B5"
	Book6.Book_id = "B6"
	Book7.Book_id = "B7"
	Book8.Book_id = "B8"
	Book9.Book_id = "B9"
	Book10.Book_id = "B10"
	Book1.Id = primitive.NewObjectID()
	Book2.Id = primitive.NewObjectID()
	Book3.Id = primitive.NewObjectID()
	Book4.Id = primitive.NewObjectID()
	Book5.Id = primitive.NewObjectID()
	Book6.Id = primitive.NewObjectID()
	Book7.Id = primitive.NewObjectID()
	Book8.Id = primitive.NewObjectID()
	Book9.Id = primitive.NewObjectID()
	Book10.Id = primitive.NewObjectID()
	Book1.Created_at = week1
	Book2.Created_at = week1
	Book3.Created_at = week1
	Book4.Created_at = week1
	Book5.Created_at = week1
	Book6.Created_at = week1
	Book7.Created_at = week1
	Book8.Created_at = week1
	Book9.Created_at = week1
	Book10.Created_at = week1
	Book1.Updated_at = week1
	Book2.Updated_at = week1
	Book3.Updated_at = week1
	Book4.Updated_at = week1
	Book5.Updated_at = now
	Book6.Updated_at = now
	Book7.Updated_at = now
	Book8.Updated_at = now
	Book9.Updated_at = week1
	Book10.Updated_at = now

	books := []interface{}{}
	books = append(books, Book1)
	books = append(books, Book2)
	books = append(books, Book3)
	books = append(books, Book4)
	books = append(books, Book5)
	books = append(books, Book6)
	books = append(books, Book7)
	books = append(books, Book8)
	books = append(books, Book9)
	books = append(books, Book10)
	BookCollection.InsertMany(ctx, books)
}

func SeedBookDetail() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	BookDetail1 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "booked",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail2 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail3 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail4 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail5 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail6 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail7 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail8 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail9 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail10 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail11 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail12 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail13 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail14 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail15 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail16 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail17 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail18 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail19 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail20 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail21 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "booked",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail22 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "booked",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail23 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail24 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail25 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail26 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail27 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail28 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail29 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail30 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail31 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail32 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail33 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail34 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail35 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail36 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail37 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail38 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail39 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail40 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail41 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail42 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail43 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail44 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail45 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail46 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail47 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail48 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail49 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail50 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "booked",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail1.Id = primitive.NewObjectID()
	BookDetail2.Id = primitive.NewObjectID()
	BookDetail3.Id = primitive.NewObjectID()
	BookDetail4.Id = primitive.NewObjectID()
	BookDetail5.Id = primitive.NewObjectID()
	BookDetail6.Id = primitive.NewObjectID()
	BookDetail7.Id = primitive.NewObjectID()
	BookDetail8.Id = primitive.NewObjectID()
	BookDetail9.Id = primitive.NewObjectID()
	BookDetail10.Id = primitive.NewObjectID()
	BookDetail11.Id = primitive.NewObjectID()
	BookDetail12.Id = primitive.NewObjectID()
	BookDetail13.Id = primitive.NewObjectID()
	BookDetail14.Id = primitive.NewObjectID()
	BookDetail15.Id = primitive.NewObjectID()
	BookDetail16.Id = primitive.NewObjectID()
	BookDetail17.Id = primitive.NewObjectID()
	BookDetail18.Id = primitive.NewObjectID()
	BookDetail19.Id = primitive.NewObjectID()
	BookDetail20.Id = primitive.NewObjectID()
	BookDetail21.Id = primitive.NewObjectID()
	BookDetail22.Id = primitive.NewObjectID()
	BookDetail23.Id = primitive.NewObjectID()
	BookDetail24.Id = primitive.NewObjectID()
	BookDetail25.Id = primitive.NewObjectID()
	BookDetail25.Id = primitive.NewObjectID()
	BookDetail26.Id = primitive.NewObjectID()
	BookDetail27.Id = primitive.NewObjectID()
	BookDetail28.Id = primitive.NewObjectID()
	BookDetail29.Id = primitive.NewObjectID()
	BookDetail30.Id = primitive.NewObjectID()
	BookDetail31.Id = primitive.NewObjectID()
	BookDetail32.Id = primitive.NewObjectID()
	BookDetail33.Id = primitive.NewObjectID()
	BookDetail34.Id = primitive.NewObjectID()
	BookDetail35.Id = primitive.NewObjectID()
	BookDetail36.Id = primitive.NewObjectID()
	BookDetail37.Id = primitive.NewObjectID()
	BookDetail38.Id = primitive.NewObjectID()
	BookDetail39.Id = primitive.NewObjectID()
	BookDetail40.Id = primitive.NewObjectID()
	BookDetail41.Id = primitive.NewObjectID()
	BookDetail42.Id = primitive.NewObjectID()
	BookDetail43.Id = primitive.NewObjectID()
	BookDetail44.Id = primitive.NewObjectID()
	BookDetail45.Id = primitive.NewObjectID()
	BookDetail46.Id = primitive.NewObjectID()
	BookDetail47.Id = primitive.NewObjectID()
	BookDetail48.Id = primitive.NewObjectID()
	BookDetail49.Id = primitive.NewObjectID()
	BookDetail50.Id = primitive.NewObjectID()

	BookDetail1.Book_detail_id = "D11"
	BookDetail2.Book_detail_id = "D12"
	BookDetail3.Book_detail_id = "D13"
	BookDetail4.Book_detail_id = "D14"
	BookDetail5.Book_detail_id = "D15"
	BookDetail6.Book_detail_id = "D21"
	BookDetail7.Book_detail_id = "D22"
	BookDetail8.Book_detail_id = "D23"
	BookDetail9.Book_detail_id = "D24"
	BookDetail10.Book_detail_id = "D25"
	BookDetail11.Book_detail_id = "D31"
	BookDetail12.Book_detail_id = "D32"
	BookDetail13.Book_detail_id = "D33"
	BookDetail14.Book_detail_id = "D34"
	BookDetail15.Book_detail_id = "D35"
	BookDetail16.Book_detail_id = "D41"
	BookDetail17.Book_detail_id = "D42"
	BookDetail18.Book_detail_id = "D43"
	BookDetail19.Book_detail_id = "D44"
	BookDetail20.Book_detail_id = "D45"
	BookDetail21.Book_detail_id = "D51"
	BookDetail22.Book_detail_id = "D52"
	BookDetail23.Book_detail_id = "D53"
	BookDetail24.Book_detail_id = "D54"
	BookDetail25.Book_detail_id = "D55"
	BookDetail26.Book_detail_id = "D61"
	BookDetail27.Book_detail_id = "D62"
	BookDetail28.Book_detail_id = "D63"
	BookDetail29.Book_detail_id = "D64"
	BookDetail30.Book_detail_id = "D65"
	BookDetail31.Book_detail_id = "D71"
	BookDetail32.Book_detail_id = "D72"
	BookDetail33.Book_detail_id = "D73"
	BookDetail34.Book_detail_id = "D74"
	BookDetail35.Book_detail_id = "D75"
	BookDetail36.Book_detail_id = "D81"
	BookDetail37.Book_detail_id = "D82"
	BookDetail38.Book_detail_id = "D83"
	BookDetail39.Book_detail_id = "D84"
	BookDetail40.Book_detail_id = "D85"
	BookDetail41.Book_detail_id = "D91"
	BookDetail42.Book_detail_id = "D92"
	BookDetail43.Book_detail_id = "D93"
	BookDetail44.Book_detail_id = "D94"
	BookDetail45.Book_detail_id = "D95"
	BookDetail46.Book_detail_id = "D101"
	BookDetail47.Book_detail_id = "D102"
	BookDetail48.Book_detail_id = "D103"
	BookDetail49.Book_detail_id = "D104"
	BookDetail50.Book_detail_id = "D105"

	BookDetail := []interface{}{}
	BookDetail = append(BookDetail, BookDetail1)
	BookDetail = append(BookDetail, BookDetail2)
	BookDetail = append(BookDetail, BookDetail3)
	BookDetail = append(BookDetail, BookDetail4)
	BookDetail = append(BookDetail, BookDetail5)
	BookDetail = append(BookDetail, BookDetail6)
	BookDetail = append(BookDetail, BookDetail7)
	BookDetail = append(BookDetail, BookDetail8)
	BookDetail = append(BookDetail, BookDetail9)
	BookDetail = append(BookDetail, BookDetail10)
	BookDetail = append(BookDetail, BookDetail11)
	BookDetail = append(BookDetail, BookDetail12)
	BookDetail = append(BookDetail, BookDetail13)
	BookDetail = append(BookDetail, BookDetail14)
	BookDetail = append(BookDetail, BookDetail15)
	BookDetail = append(BookDetail, BookDetail16)
	BookDetail = append(BookDetail, BookDetail17)
	BookDetail = append(BookDetail, BookDetail18)
	BookDetail = append(BookDetail, BookDetail19)
	BookDetail = append(BookDetail, BookDetail20)
	BookDetail = append(BookDetail, BookDetail21)
	BookDetail = append(BookDetail, BookDetail22)
	BookDetail = append(BookDetail, BookDetail23)
	BookDetail = append(BookDetail, BookDetail24)
	BookDetail = append(BookDetail, BookDetail25)
	BookDetail = append(BookDetail, BookDetail26)
	BookDetail = append(BookDetail, BookDetail27)
	BookDetail = append(BookDetail, BookDetail28)
	BookDetail = append(BookDetail, BookDetail29)
	BookDetail = append(BookDetail, BookDetail30)
	BookDetail = append(BookDetail, BookDetail31)
	BookDetail = append(BookDetail, BookDetail32)
	BookDetail = append(BookDetail, BookDetail33)
	BookDetail = append(BookDetail, BookDetail34)
	BookDetail = append(BookDetail, BookDetail35)
	BookDetail = append(BookDetail, BookDetail36)
	BookDetail = append(BookDetail, BookDetail37)
	BookDetail = append(BookDetail, BookDetail38)
	BookDetail = append(BookDetail, BookDetail39)
	BookDetail = append(BookDetail, BookDetail40)
	BookDetail = append(BookDetail, BookDetail41)
	BookDetail = append(BookDetail, BookDetail42)
	BookDetail = append(BookDetail, BookDetail43)
	BookDetail = append(BookDetail, BookDetail44)
	BookDetail = append(BookDetail, BookDetail45)
	BookDetail = append(BookDetail, BookDetail46)
	BookDetail = append(BookDetail, BookDetail47)
	BookDetail = append(BookDetail, BookDetail48)
	BookDetail = append(BookDetail, BookDetail49)
	BookDetail = append(BookDetail, BookDetail50)
	BookDetailCollection.InsertMany(ctx, BookDetail)
}

func SeedBookRent() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookRent1 := models.BookRentModel{
		Book_id:      "B1",
		User_id:      "U02",
		Reserve_date: now,
	}
	BookRent2 := models.BookRentModel{
		Book_id:      "B5",
		User_id:      "U03",
		Reserve_date: now,
	}
	BookRent3 := models.BookRentModel{
		Book_id:      "B5",
		User_id:      "U04",
		Reserve_date: now,
	}
	BookRent4 := models.BookRentModel{
		Book_id:      "B10",
		User_id:      "U05",
		Reserve_date: now,
	}
	BookRent1.Book_detail_id = "D11"
	BookRent2.Book_detail_id = "D51"
	BookRent3.Book_detail_id = "D52"
	BookRent4.Book_detail_id = "D105"
	BookRent1.Book_rent_id = "R1"
	BookRent2.Book_rent_id = "R2"
	BookRent3.Book_rent_id = "R3"
	BookRent4.Book_rent_id = "R4"

	BookRent1.Id = primitive.NewObjectID()
	BookRent2.Id = primitive.NewObjectID()
	BookRent3.Id = primitive.NewObjectID()
	BookRent4.Id = primitive.NewObjectID()

	BookRent := []interface{}{
		BookRent1,
		BookRent2,
		BookRent3,
		BookRent4,
	}
	BookRentCollection.InsertMany(ctx, BookRent)
}

func SeedBookBorrow() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookBorrow1 := models.BookBorrowModel{
		User_id:        "U02",
		Book_id:        "B6",
		Book_detail_id: "B61",
	}
	BookBorrow2 := models.BookBorrowModel{
		User_id:        "U03",
		Book_id:        "B6",
		Book_detail_id: "B62",
	}
	BookBorrow3 := models.BookBorrowModel{
		User_id:        "U04",
		Book_id:        "B7",
		Book_detail_id: "D75",
	}
	BookBorrow4 := models.BookBorrowModel{
		User_id:        "U05",
		Book_id:        "B8",
		Book_detail_id: "D81",
	}
	BookBorrow1.Id = primitive.NewObjectID()
	BookBorrow2.Id = primitive.NewObjectID()
	BookBorrow3.Id = primitive.NewObjectID()
	BookBorrow4.Id = primitive.NewObjectID()
	BookBorrow1.Book_hire_id = "BR01"
	BookBorrow2.Book_hire_id = "BR02"
	BookBorrow3.Book_hire_id = "BR03"
	BookBorrow4.Book_hire_id = "BR04"
	BookBorrows := []interface{}{
		BookBorrow1,
		BookBorrow2,
		BookBorrow3,
		BookBorrow4,
	}
	BookBorrowCollection.InsertMany(ctx, BookBorrows)
}

func SeedHistory() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	History1 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B3",
		Book_detail_id: "B31",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History2 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B5",
		Book_detail_id: "B51",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History3 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B10",
		Book_detail_id: "B101",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History4 := models.HistoryModel{
		User_id:        "U05",
		Book_id:        "B10",
		Book_detail_id: "B101",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History5 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B10",
		Book_detail_id: "B102",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History6 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B10",
		Book_detail_id: "B103",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History7 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B10",
		Book_detail_id: "B104",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History8 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B4",
		Book_detail_id: "B41",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History9 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B5",
		Book_detail_id: "B53",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History10 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B6",
		Book_detail_id: "B64",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History11 := models.HistoryModel{
		User_id:        "U05",
		Book_id:        "B7",
		Book_detail_id: "B71",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History12 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B9",
		Book_detail_id: "B93",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History1.Id = primitive.NewObjectID()
	History2.Id = primitive.NewObjectID()
	History3.Id = primitive.NewObjectID()
	History4.Id = primitive.NewObjectID()
	History5.Id = primitive.NewObjectID()
	History6.Id = primitive.NewObjectID()
	History7.Id = primitive.NewObjectID()
	History8.Id = primitive.NewObjectID()
	History9.Id = primitive.NewObjectID()
	History10.Id = primitive.NewObjectID()
	History11.Id = primitive.NewObjectID()
	History12.Id = primitive.NewObjectID()
	History1.History_id = "H01"
	History1.History_id = "H02"
	History1.History_id = "H03"
	History1.History_id = "H04"
	History1.History_id = "H05"
	History1.History_id = "H06"
	History1.History_id = "H07"
	History1.History_id = "H08"
	History1.History_id = "H09"
	History1.History_id = "H10"
	History1.History_id = "H11"
	History1.History_id = "H12"
	Histories := []interface{}{
		History1,
		History2,
		History3,
		History4,
		History5,
		History6,
		History7,
		History8,
		History9,
		History10,
		History11,
		History12,
	}
	HistoryCollection.InsertMany(ctx, Histories)
}

func SeedType() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	Type1 := models.BookTypeModel{
		TypeId:   "1",
		TypeName: "Manga",
	}
	Type2 := models.BookTypeModel{
		TypeId:   "2",
		TypeName: "Education",
	}
	Type3 := models.BookTypeModel{
		TypeId:   "3",
		TypeName: "IT",
	}
	Type1.Id = primitive.NewObjectID()
	Type2.Id = primitive.NewObjectID()
	Type3.Id = primitive.NewObjectID()
	Types := []interface{}{
		Type1,
		Type2,
		Type3,
	}
	TypeCollections.InsertMany(ctx, Types)
}
