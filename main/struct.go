package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var albert Customer
	albert.Name = "Albert"
	albert.Address = "Indoneisia"
	albert.Age = 27

	fmt.Println(albert.Name)
	fmt.Println(albert.Address)
	fmt.Println(albert.Age)

	pras := Customer{
		Name:    "Pras",
		Address: "Jawa",
		Age:     30,
	}

	fmt.Println(pras.Name)
	fmt.Println(pras.Address)
	fmt.Println(pras.Age)

	tomo := Customer{"Tomo", "Sunda", 40}

	fmt.Println(tomo.Name)
	fmt.Println(tomo.Address)
	fmt.Println(tomo.Age)

}
