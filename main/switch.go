package main

import "fmt"

func main() {
	name := "sad"

	switch name {
	case "Albert":
		fmt.Println(name)
		fmt.Println(name)
	default:
		fmt.Println("Kamu Bukan Albert")
		fmt.Println("Kamu Bukan Albert")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("true")
		fmt.Println(length)
	case false:
		fmt.Println("false")
		fmt.Println(length)
	}

	length := len(name)
	switch {
	case length > 7:
		fmt.Println("true")
		fmt.Println(length)
	default:
		fmt.Println("false")
		fmt.Println(length)
	}

}
