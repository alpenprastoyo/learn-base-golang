package main

import "fmt"

func main() {
	name := "Kap"

	if name == "Kapit" {
		fmt.Println("Hello, Kapit!")
	} else if name == "Kapist" {
		fmt.Println("Hello, Kapist!")
	} else {
		fmt.Println("kamu Bukan, Kapit!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Aman")
	}
}
