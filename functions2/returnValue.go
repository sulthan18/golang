package main

import "fmt"

func sayGreetings(name string) string {
	return "hello " + name
}

func main() {
	result := sayGreetings("sulthan")
	fmt.Println(result)
}
