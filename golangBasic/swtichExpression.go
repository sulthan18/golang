package main

import "fmt"

func main() {
	name := "sulthan"

	switch name {
	case "sulthan":
		fmt.Println("hello sulthan")
	case "jaka":
		fmt.Println("hello jaka")
	default:
		fmt.Println("hi,boleh kenalan?")
	}
}
