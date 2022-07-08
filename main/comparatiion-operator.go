package main

import "fmt"

func main() {
	var name1 = "Kapit"
	var name2 = "Kapit"

	var result bool = name1 == name2
	fmt.Println(result)

	var name3 = "Kapit"
	var name4 = "Kapit Pramudito"

	var result2 bool = name3 > name4
	fmt.Println(result2)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
