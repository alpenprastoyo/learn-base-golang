package main

import "fmt"

func main() {
	var months = [...]string{
		"january",
		"february",
		"march",
		"april",
		"mei",
		"june",
		"july",
		"august",
		"september",
		"oktober",
		"november",
		"december",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"

	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Libur")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Kapit"
	newSlice[1] = "Pramudito"

	fmt.Println(newSlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
