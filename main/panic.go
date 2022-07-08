package main

import "fmt"

func main(){
	runApp(true);
}

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Error dengan mesasge:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}