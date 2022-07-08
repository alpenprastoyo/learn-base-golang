package main

import "fmt"

func main() {
	var name string
	var name2 = "Pramudito Wasi"

	name = "coba string"

	fmt.Println(name)

	name = "Kapit Pramudito"

	fmt.Println(name)

	fmt.Println("nama = ", name2)

	var age = 69

	fmt.Println("umur = ", age)

	lokasi := "UNS"

	fmt.Println("Lokasi = ", lokasi)

	var (
		kecamatan = "Jebres"
		kota      = "Surakarta"
	)

	fmt.Println("Kecamatan = ", kecamatan)
	fmt.Println("Kota = ", kota)

}
