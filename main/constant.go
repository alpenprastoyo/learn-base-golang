package main

import "fmt"

func main() {

	//single const
	const firstname string = "Kapit"
	const lastname string = "Pramudito"
	const age = 30

	//multiple const
	const (
		kecamatan = "jebres"
		kota      = "surakarta"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(age)
	fmt.Println(kecamatan)
	fmt.Println(kota)

}
