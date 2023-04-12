package main

import "fmt"

// func getFullName() (string, string) {
// 	return "sulthan", "hendriansyah"
// }

// func main() {
// 	firstName, lastName := getFullName()
// 	fmt.Println(firstName, lastName)

func sayHello() (string, string) {
	return "Hello", "Michelle"
}

func main() {
	hello, _ := sayHello()
	fmt.Println(hello)
}
