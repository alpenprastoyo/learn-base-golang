package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println(total)

	result, err := calculation(10, 2, "+")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = calculation(10, 2, "-")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = calculation(10, 2, "*")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = calculation(10, 2, "/")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = calculation(10, 2, "=")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sum(scores []int) int {
	total := 0
	for _, value := range scores {
		total += value
	}

	return total
}

func calculation(numberOne float32, numberTwo float32, operation string) (float32, error) {
	var result float32
	var errorResult error

	switch operation {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		result = numberOne / numberTwo
	default:
		errorResult = errors.New("Unknown Opertaion")
	}

	return result, errorResult
}
