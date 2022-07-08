package main

import "fmt"

func main() {
	firstName, lastName := getFullName()

	fmt.Printf(firstName)
	fmt.Println(lastName)

	namaPertama, _ := getFullName()

	fmt.Println(namaPertama)
}

func getFullName() (string, string) {
	return "Abert", "Tomo"
}
