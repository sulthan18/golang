package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("perulangan ke", counter)
	// }

	names := []string{"sulthan", "rafi", "hendriansyah"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

}
