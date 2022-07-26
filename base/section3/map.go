package main

import "fmt"

func main() {
	//map[type data of key]type data of value
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["PHP"] = 9
	myMap["go"] = 10

	fmt.Println(myMap["ruby"])

	myMap2 := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"Javascript": "hype",
	}

	fmt.Println(myMap2)

	//loop for map

	for key, value := range myMap2 {
		fmt.Println("Key : ", key, "value: ", value)
	}

	delete(myMap2, "ruby")

	fmt.Println("=========")

	for key, value := range myMap2 {
		fmt.Println("Key : ", key, "value: ", value)
	}

	value, isAvailable := myMap2["python"]

	fmt.Println(isAvailable)
	fmt.Println(value)

}
