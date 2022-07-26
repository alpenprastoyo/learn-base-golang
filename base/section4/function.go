package main

import "fmt"

func main() {
	sentence := printMyResult("Go")
	fmt.Println(sentence)

	result := add(5, 9)
	fmt.Println(result)

	area, perimeter := rectangular(10, 5)
	fmt.Println("Area of rectangular : ", area)
	fmt.Println("Perimeter of rectangular: ", perimeter)

	areaRectangle, perimeterRectangle := rectangle(10)
	fmt.Println("Area of rectangle : ", areaRectangle)
	fmt.Println("Perimeter of rectangle: ", perimeterRectangle)

}

func printMyResult(language string) string {
	sentence := "Saya sedang belajar " + language
	return sentence
}

func add(numberOne int, numberTwo int) int {
	result := numberOne + numberTwo
	return result
}

func rectangular(length int, width int) (int, int) {
	area := length * width
	perimeter := 2*length + 2*width
	return area, perimeter
}

func rectangle(side int) (area int, perimeter int) {
	area = side * side
	perimeter = 4 * side
	return
}
