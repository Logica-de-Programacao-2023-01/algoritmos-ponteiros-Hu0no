package main

type Book struct {
	Title  string
	Author string
	Price  float64
}

func PriceWithDiscount(ptr *Book) {
	ptr.Price = 0.9 * ptr.Price
}
