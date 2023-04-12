package main

import (
	"fmt"
)

func main() {
	// var nilai32 int32 = 129387
	// var nilai64 int64 = int64(nilai32)

	// var nilai16 int16 = int16(nilai32)

	// fmt.Println(nilai16)
	// fmt.Println(nilai32)
	// fmt.Println(nilai64)

	var (
		nilai32 int32 = 12938
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)

		name = "sulthan rafi hendriansyah"
		s    = string(name[0])
	)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(name)
	fmt.Println(s)

}
