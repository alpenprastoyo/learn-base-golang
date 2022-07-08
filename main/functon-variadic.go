package main

import "fmt"

func main() {

	total := sumAll(1, 2, 3, 4)

	fmt.Println(total)

	numbers := []int{10, 20, 30, 40, 50}
	total2 := sumAll(numbers...)
	fmt.Println(total2)

}

func sumAll(numbers ...int) int {

	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
