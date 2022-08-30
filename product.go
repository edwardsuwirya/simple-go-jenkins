package main

type Product struct {
	Id           string
	ProductName  string `db:"product_name"`
	CategoryId   string `db:"category_id"`
	CategoryName string `db:"category_name"`
}
