package main

import "fmt"

func main() {
	age := 9
	var grade string

	if age > 10 {
		fmt.Println("Your age more than 10 years")
	} else {
		fmt.Println("Your age less than 10 years")
	}

	score := 65

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else if score < 80 {
		grade = "B"
	} else {
		grade = "A"
	}

	fmt.Println(grade)
}
