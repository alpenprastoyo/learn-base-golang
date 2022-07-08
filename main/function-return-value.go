package main

import "fmt"

func helloThere(name string) string {
	if name == "" {
		return "Hello Anonimous"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := helloThere("")
	fmt.Println(result)
}
