package main

import "fmt"

func main() {
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "PS 4")
	gamingConsole = append(gamingConsole, "PS 3")
	gamingConsole = append(gamingConsole, "PS 2")
	gamingConsole = append(gamingConsole, "PS 1")

	game := []string{
		"COD 1",
		"COD 2",
		"COD 3",
		"COD 4",
	}

	fmt.Println(gamingConsole)
	fmt.Println(game)

}
