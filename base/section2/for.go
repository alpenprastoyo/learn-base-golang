package main

import "fmt"

func main() {

	for i := 1; i < 10000; i++ {
		fmt.Println("I'm learn Go :", i)
	}

	g := 1
	for g <= 10000 {
		fmt.Println("I'm learn Go :", g)
		g++
	}

	title := "Golang the nest language"

	for index, letter := range title {
		fmt.Println("index :", index, " letter: ", string(letter))
	}
}
