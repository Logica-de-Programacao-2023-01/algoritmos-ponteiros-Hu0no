package main

type Book struct {
	Title  string
	Author string
}

func UpdateBookInfo(ptr *Book) {
	if ptr.Author == "Anonymous" {
		ptr.Title = "Unknown"
	}
}
