package main

import "fmt"

func main() {

	a, b, c := getFullName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Abert"
	middleName = "Pras"
	lastName = "Tomo"

	return
}
