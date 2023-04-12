package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Sulthan Rafi Hendriansyah",
		"country": "Indonesia",
	}
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "belajar go-lang dasar"
	book["author"] = "Michelle Wijaya"
	book["co-author"] = "Sulthan Rafi Hendriansyah"

	delete(book, "author")

	fmt.Println(book)

}
