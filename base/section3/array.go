package main

import "fmt"

func main() {
	var languages [5]string
	languages2 := [5]string{"Go", "Ruby", "Pythin", "C++", "Php"}

	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "Php"
	languages[3] = "V++"
	languages[4] = "Python"
	fmt.Println(languages2)

	length := len(languages2)

	fmt.Println(length)

	languages3 := [...]string{
		"Go",
		"Ruby",
		"Pythin",
		"C++",
		"Php",
	}

	for index, lang := range languages3 {
		fmt.Println("index : ", index, " Languages: ", lang)
	}

}
