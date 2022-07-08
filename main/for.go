package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("counter: ", counter)
		counter++
	}

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("counter: ", counter)
	}

	names := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(names); i++ {
		fmt.Println("data: ", names[i])
	}

	for index, name := range names {
		fmt.Println("index: ", index, "=", name)
	}

	for name := range names {
		fmt.Println(name)
	}

	books := make(map[string]string)
	books["title"] = "Sapiens"
	books["genre"] = "Historical"

	for key, value := range books {
		fmt.Println("key: ", key, "=", value)
	}
}
