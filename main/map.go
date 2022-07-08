package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Kapit",
		"address": "Surakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Sapiens"
	book["genre"] = "Historical"
	book["salah"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "salah")

	fmt.Println(book)
	fmt.Println(len(book))

}
