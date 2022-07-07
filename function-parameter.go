package main

import "fmt"

func main() {
	name("Albert", "Tomo")
}

func name(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}
