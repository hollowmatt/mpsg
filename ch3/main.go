package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(bookworms)
	bookCount := booksCount(bookworms)
	fmt.Println(bookCount)
	commonBooks := findCommonBooks(bookworms)
	fmt.Println(commonBooks)
}
