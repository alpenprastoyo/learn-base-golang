package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your host")
	var password *string = flag.String("password", "root", "Put your password here")
	var number *int = flag.Int("number", 100, "Put your password here")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("Password :", *password)
	fmt.Println("Number :", *number)

}
