package main

import "fmt"

func main() {
	title := "Golang the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("indext = ", index, " letter :", string(letter))
		}
	}

	title = "Golang the best language"

	for index, letter := range title {
		character := string(letter)

		if character == "a" || character == "i" || character == "u" || character == "e" || character == "o" {
			fmt.Println("indext = ", index, " letter :", string(letter))
		}

		switch character {
		case "a", "i", "u", "e", "o":
			fmt.Println("indext = ", index, " letter :", string(letter))
		}
	}
}
