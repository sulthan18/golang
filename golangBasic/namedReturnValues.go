package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "michelle"
	middleName = "ayze"
	lastName = "ken"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
