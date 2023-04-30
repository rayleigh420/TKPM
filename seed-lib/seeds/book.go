package seeds

import (
	"context"
	"time"

	"github.com/baguette/seed-lib/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var now,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
func SeedBook(){
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	Book1 := models.BookModel{
		Name: "Không Diệt Không Sinh Đừng Sợ Hãi",
		Publisher: "Saigon Books",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Thích Nhất Hạnh",
		Borrowed_quantity: 12,
		Amount: 50,
		Type_id: "2",
	}
	Book2 := models.BookModel{
		Name: "Muôn Kiếp Nhân Sinh - Many Times, Many Lives",
		Publisher: "NXB Tổng Hợp TPHCM",
		Yearpublished: 2020,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Nguyên Phong",
		Borrowed_quantity: 120,
		Amount: 50,
		Type_id: "2",
	}
	Book3 := models.BookModel{
		Name: "Muốn An Được An",
		Publisher: "Thái Hà",
		Yearpublished: 2021,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Thích Nhất Hạnh",
		Borrowed_quantity: 20,
		Amount: 50,
		Type_id: "2",
	}
	Book4 := models.BookModel{
		Name: "Tĩnh Lặng - Sức Mạnh Tĩnh Lặng Trong Thế Giới Huyền Ảo",
		Publisher: "Thái Hà",
		Yearpublished: 2020,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Thích Nhất Hạnh",
		Borrowed_quantity: 42,
		Amount: 50,
		Type_id: "2",
	}
	Book5 := models.BookModel{
		Name: "Chainsaw Man - Tập 11",
		Publisher: "NXB Trẻ",
		Yearpublished: 2020,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Tatsuki Fujimoto",
		Borrowed_quantity: 512,
		Amount: 50,
		Type_id: "1",
	}
	Book6 := models.BookModel{
		Name: "Khu Rừng Đom Đóm",
		Publisher: "NXB Trẻ",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Yuki Midorikawa",
		Borrowed_quantity: 162,
		Amount: 50,
		Type_id: "1",
	}
	Book7 := models.BookModel{
		Name: "Cô Gái Nơi Xứ Ngoài",
		Publisher: "IPM",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Nagabe",
		Borrowed_quantity: 127,
		Amount: 50,
		Type_id: "1",
	}
	Book8 := models.BookModel{
		Name: "Cô Gái Nơi Xứ Ngoài - Tập 2",
		Publisher: "IPM",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Nagabe",
		Borrowed_quantity: 182,
		Amount: 50,
		Type_id: "1",
	}
	Book9 := models.BookModel{
		Name: "Blockchain - Bản Chất Của Blockchain",
		Publisher: "1980 Books",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Mark Gates",
		Borrowed_quantity: 192,
		Amount: 50,
		Type_id: "3",
	}
	Book10 := models.BookModel{
		Name: "Object Oriented Programming",
		Publisher: "TKLong",
		Yearpublished: 2000,
		Book_img: "https://bit.ly/41PTfoj",
		Author: "Nguyen Van A",
		Borrowed_quantity: 1012,
		Amount: 50,
		Type_id: "3",
	}
	
	// Book1.Book_id = primitive.NewObjectID().Hex()
	// Book2.Book_id = primitive.NewObjectID().Hex()
	// Book3.Book_id = primitive.NewObjectID().Hex()
	// Book4.Book_id = primitive.NewObjectID().Hex()
	// Book5.Book_id = primitive.NewObjectID().Hex()
	// Book6.Book_id = primitive.NewObjectID().Hex()
	// Book7.Book_id = primitive.NewObjectID().Hex()
	// Book8.Book_id = primitive.NewObjectID().Hex()
	// Book9.Book_id = primitive.NewObjectID().Hex()
	// Book10.Book_id= primitive.NewObjectID().Hex()
	// Book1.Id,_ = primitive.ObjectIDFromHex(Book1.Book_id)
	// Book2.Id,_ = primitive.ObjectIDFromHex(Book2.Book_id)
	// Book3.Id,_ = primitive.ObjectIDFromHex(Book3.Book_id)
	// Book4.Id,_ = primitive.ObjectIDFromHex(Book4.Book_id)
	// Book5.Id,_ = primitive.ObjectIDFromHex(Book5.Book_id)
	// Book6.Id,_ = primitive.ObjectIDFromHex(Book6.Book_id)
	// Book7.Id,_ = primitive.ObjectIDFromHex(Book7.Book_id)
	// Book8.Id,_ = primitive.ObjectIDFromHex(Book8.Book_id)
	// Book9.Id,_ = primitive.ObjectIDFromHex(Book9.Book_id)
	// Book10.Id,_ = primitive.ObjectIDFromHex(Book10.Book_id)

	Book1.Book_id = "8bNiueHjc9cU"
	Book2.Book_id = "TXefnQqaL3C3"
	Book3.Book_id = "faHvUfqdd87I"
	Book4.Book_id = "xgroablsJhst"
	Book5.Book_id = "uvWaWqqFrgn1"
	Book6.Book_id = "5rgxmEyyuRVR"
	Book7.Book_id = "8F4W93lFMJSh"
	Book8.Book_id = "NvvnZH7xqLBi"
	Book9.Book_id = "LgvTiRr6hDLN"
	Book10.Book_id= "RbPuAnhCPQsw"
	Book1.Id= primitive.NewObjectID()
	Book2.Id= primitive.NewObjectID()
	Book3.Id = primitive.NewObjectID()
	Book4.Id = primitive.NewObjectID()
	Book5.Id = primitive.NewObjectID()
	Book6.Id = primitive.NewObjectID()
	Book7.Id = primitive.NewObjectID()
	Book8.Id = primitive.NewObjectID()
	Book9.Id = primitive.NewObjectID()
	Book10.Id= primitive.NewObjectID()

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
	BookCollection.InsertMany(ctx,books)
}

func SeedBookDetail(){
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	
	BookDetail1 := models.BookDetailModel{
		Book_id: "TXefnQqaL3C3",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail2 := models.BookDetailModel{
		Book_id: "TXefnQqaL3C3",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail3 := models.BookDetailModel{
		Book_id: "uvWaWqqFrgn1",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail4 := models.BookDetailModel{
		Book_id: "uvWaWqqFrgn1",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail5 := models.BookDetailModel{
		Book_id: "uvWaWqqFrgn1",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail6 := models.BookDetailModel{
		Book_id: "uvWaWqqFrgn1",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail7 := models.BookDetailModel{
		Book_id: "RbPuAnhCPQsw",
		Location:"ABC123",
		Status:"ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail8 := models.BookDetailModel{
		Book_id: "RbPuAnhCPQsw",
		Location:"ABC123",
		Status:"borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail9 := models.BookDetailModel{
		Book_id: "RbPuAnhCPQsw",
		Location:"ABC123",
		Status:"borrowed",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail10 := models.BookDetailModel{
		Book_id: "RbPuAnhCPQsw",
		Location:"ABC123",
		Status:"borrowed",
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
	
	BookDetail1.Book_detail_id = "eGmZrEEM7PFb"
	BookDetail2.Book_detail_id = "9vLrivXOlcHr"
	BookDetail3.Book_detail_id = "APiac1eTN8Ao"
	BookDetail4.Book_detail_id = "qsUEcPXQIFdx"
	BookDetail5.Book_detail_id = "ahDyuPOcM83G"
	BookDetail6.Book_detail_id = "BvfFnLVSfve8"
	BookDetail7.Book_detail_id = "VM9rSlejHqBi"
	BookDetail8.Book_detail_id = "97CJjBOxVRzS"
	BookDetail9.Book_detail_id = "OarWZhIKctGr"
	BookDetail10.Book_detail_id = "k7MAMwNrtPdh"

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
	BookDetailCollection.InsertMany(ctx,BookDetail)
}

func SeedBookRent(){
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookRent1 := models.BookRentModel{
		Book_id: "RbPuAnhCPQsw",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent2 := models.BookRentModel{
		Book_id: "RbPuAnhCPQsw",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent3 := models.BookRentModel{
		Book_id: "uvWaWqqFrgn1",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent4 := models.BookRentModel{
		Book_id: "uvWaWqqFrgn1",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent5 := models.BookRentModel{
		Book_id: "TXefnQqaL3C3",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent6 := models.BookRentModel{
		Book_id: "TXefnQqaL3C3",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent7 := models.BookRentModel{
		Book_id: "TXefnQqaL3C3",
		User_id:"6ZrgtsjaKQi7",
		Reserve_date: now,
	}
	BookRent8 := models.BookRentModel{
		Book_id: "uvWaWqqFrgn1",
		User_id:"UtkvLPlXNzm4",
		Reserve_date: now,
	}
	BookRent9 := models.BookRentModel{
		Book_id: "uvWaWqqFrgn1",
		User_id:"6clomWjyrM1n",
		Reserve_date: now,
	}
	BookRent10 := models.BookRentModel{
		Book_id: "uvWaWqqFrgn1",
		User_id:"6clomWjyrM1n",
		Reserve_date: now,
	}
	BookRent1.Book_rent_id = "SQMsZ3y2NFQb"
	BookRent2.Book_rent_id = "xduD5RQSWUe5"
	BookRent3.Book_rent_id = "12CWC2wJz466"
	BookRent4.Book_rent_id = "4XZOX5ArmClm"
	BookRent5.Book_rent_id = "jRZl1E8XTkz6"
	BookRent6.Book_rent_id = "l9Uj51rOVW9S"
	BookRent7.Book_rent_id = "CepcFb3BvLOR"
	BookRent8.Book_rent_id = "ph87nmpFQZO8"
	BookRent9.Book_rent_id = "hpzbhd2153KX"
	BookRent10.Book_rent_id = "71rrYohEhtaP"

	BookRent1.Id = primitive.NewObjectID()
	BookRent2.Id = primitive.NewObjectID()
	BookRent3.Id = primitive.NewObjectID()
	BookRent4.Id = primitive.NewObjectID()
	BookRent5.Id = primitive.NewObjectID()
	BookRent6.Id = primitive.NewObjectID()
	BookRent7.Id = primitive.NewObjectID()
	BookRent8.Id = primitive.NewObjectID()
	BookRent9.Id = primitive.NewObjectID()
	BookRent10.Id = primitive.NewObjectID()
	BookRent := []interface{}{
		BookRent1,
		BookRent2,
		BookRent3,
		BookRent4,
		BookRent5,
		BookRent6,
		BookRent7,
		BookRent8,
		BookRent9,
		BookRent10,
	}
	BookRentCollection.InsertMany(ctx,BookRent)
}

func SeedBookBorrow(){
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookBorrow1 := models.BookBorrowModel{
		User_id: "6ZrgtsjaKQi7",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "k7MAMwNrtPdh",
	}
	BookBorrow2 := models.BookBorrowModel{
		User_id: "6ZrgtsjaKQi7",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "OarWZhIKctGr",
	}
	BookBorrow3 := models.BookBorrowModel{
		User_id: "6ZrgtsjaKQi7",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "97CJjBOxVRzS",
	}
	BookBorrow1.Id = primitive.NewObjectID()
	BookBorrow2.Id = primitive.NewObjectID()
	BookBorrow3.Id = primitive.NewObjectID()
	BookBorrow1.Book_hire_id = "tIcyuvHKb55m"	
	BookBorrow2.Book_hire_id = "HwNmRq1WhLpH"	
	BookBorrow3.Book_hire_id = "HtIFesemOWKC"	
	BookBorrows := []interface{}{
		BookBorrow1,
		BookBorrow2,
		BookBorrow3,
	}
	BookBorrowCollection.InsertMany(ctx,BookBorrows)
}

func SeedHistory(){
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var week1,_ = time.Parse(time.RFC3339,time.Now().Add(-7*24*time.Hour).Format(time.RFC3339))
	History1 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "TXefnQqaL3C3",
		Book_detail_id: "eGmZrEEM7PFb",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History2 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "TXefnQqaL3C3",
		Book_detail_id: "9vLrivXOlcHr",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History3 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "uvWaWqqFrgn1",
		Book_detail_id: "uvWaWqqFrgn1",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History4 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "uvWaWqqFrgn1",
		Book_detail_id: "qsUEcPXQIFdx",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History5 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "uvWaWqqFrgn1",
		Book_detail_id: "ahDyuPOcM83G",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History6 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "TXefnQqaL3C3",
		Book_detail_id: "eGmZrEEM7PFb",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History7 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "TXefnQqaL3C3",
		Book_detail_id: "eGmZrEEM7PFb",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History8 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "TXefnQqaL3C3",
		Book_detail_id: "eGmZrEEM7PFb",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History9 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "OarWZhIKctGr",
		Status: "returned",
		Date_borrowed: week1,
		Date_return: now,
	}
	History10 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "k7MAMwNrtPdh",
		Status: "borrowing",
		Date_borrowed: week1,
	}
	History11 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "OarWZhIKctGr",
		Status: "borrowing",
		Date_borrowed: week1,
	}
	History12 := models.HistoryModel{
		User_id: "6clomWjyrM1n",
		Book_id: "RbPuAnhCPQsw",
		Book_detail_id: "97CJjBOxVRzS",
		Status: "borrowing",
		Date_borrowed: week1,
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
	History1.History_id = "X1ZKA7O3OvlR"
	History1.History_id = "ufhFpBq4F9gG"
	History1.History_id = "B2L5x9eIC4Hu"
	History1.History_id = "J2rUrmzWVcU2"
	History1.History_id = "NOM2cjsVHMfU"
	History1.History_id = "udfEqGOigZst"
	History1.History_id = "mNEx1IfmEFJT"
	History1.History_id = "s5BosMw3hMhB"
	History1.History_id = "sYK1rOna1OsI"
	History1.History_id = "4kp5QWqt5dQ2"
	History1.History_id = "YBVtvc4K5PT7"
	History1.History_id = "Ds8Y44chIZJq"
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
	HistoryCollection.InsertMany(ctx,Histories)
}