package main

import "fmt"

func main() {
	type biodata string

	var (
		name biodata = "Sulthan Rafi Hendriansyah"
	)
	fmt.Println(name)
	fmt.Println(biodata("jessica adeline"))
}
