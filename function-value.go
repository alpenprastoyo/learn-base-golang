package main

import "fmt"

func main() {
	hello := helloFunc

	result := hello("albert")

	fmt.Println(result)
}

func helloFunc(name string) string {
	return "Hello, " + name
}
