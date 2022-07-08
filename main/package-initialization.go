package main

import (
	"fmt"

	"tutorial/belajar-golang-dasar/main/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
