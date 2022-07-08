package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHi(name string){
	fmt.Println("hello ", customer.Name, " My name is ", name)
}

func main(){
	customer := Customer{Name: "Albert"}
	customer.sayHi("Tomo")
}

