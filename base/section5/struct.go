package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {

	user1 := User{}
	user1.ID = 1
	user1.FirstName = "Albert"
	user1.LastName = "Tomo"
	user1.Email = "alberttomo@gmail.com"
	user1.IsActive = true
	fmt.Println(user1)

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Revi"
	user2.LastName = "Shirai"
	user2.Email = "ReviShirai@gmail.com"
	user2.IsActive = true
	fmt.Println(user2)

	user3 := User{
		ID:        3,
		FirstName: "Makoto",
		LastName:  "Shinkai",
		Email:     "makotoshikai@gmail.com",
		IsActive:  true,
	}
	fmt.Println(user3)

	name := displayUser(user3)

	fmt.Println(name)

}

func displayUser(user User) string {
	result := fmt.Sprintf("name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)

	return result
}
