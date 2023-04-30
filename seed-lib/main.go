package main

import (
	// "fmt"

	"github.com/baguette/seed-lib/seeds"
	// gonanoid "github.com/matoous/go-nanoid/v2"
)



func main() {
	seeds.SeedUser()
	seeds.SeedBook()
	seeds.SeedBookDetail()
	seeds.SeedBookRent()
	seeds.SeedBookBorrow()
	seeds.SeedHistory()
	seeds.SeedType()
	// for i:=1;i<=100;i++ {
	// 	id, _ := gonanoid.Generate("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 12)
	// 	fmt.Println(id)
	// }
}
