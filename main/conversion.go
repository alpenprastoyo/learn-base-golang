package main

import "fmt"

func main() {
	var nilai32 int32 = 200
	nilai64 := int64(nilai32)
	nilai8 := int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "Kapit"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
