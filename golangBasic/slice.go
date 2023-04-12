package main

import "fmt"

func main() {
	// names := [...]string{"sulthan", "rafi", "hendriansyah", "michelle", "jaey", "yin"}
	// slice := names[4:6]

	// fmt.Println(names)
	// fmt.Println(slice)

	days := [...]string{"monday", "thursday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "new saturday"
	daySlice1[1] = "new sunday"
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "sulthan"
	newSlice[1] = "hendriansyah"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	thisIsArray := [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
