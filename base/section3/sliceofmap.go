package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "agung", "score": "A"},
		{"name": "hapsoro", "score": "B"},
		{"name": "line", "score": "C"},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println("name: ", student["name"], " score : ", student["score"])
	}
}
