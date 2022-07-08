package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len("Eko"))
	fmt.Println("Eko Setiawan"[0])
	fmt.Println("Eko Setiawan Pramudito"[1])

	fmt.Println(strings.Contains("Albert Wito", "Albert"))
	fmt.Println(strings.Contains("Albest Wito", "tomo"))
}
