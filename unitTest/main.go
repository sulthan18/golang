// case 1
// package main

// import "fmt"

// func Hello() {
// 	fmt.Println("welcome back,user!")
// }

// func main() {
// 	Hello()
// }

// case 2
package hello

import "testing"

func Hello() string {
	return ("hello,user!")
}

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
