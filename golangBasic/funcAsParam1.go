package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fmt.Println("hello ", filter(name))
// }

// func spamFilter(name string) string {
// 	if name == "kucing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
