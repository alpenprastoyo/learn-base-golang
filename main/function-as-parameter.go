package main

import (
	"fmt"
)

type Filter func(string) string

func main() {

	filter := spamFilter

	helloWithFilter("albert", filter)

}

func helloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "....."
	} else {
		return name
	}
}
