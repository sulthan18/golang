package main

import "fmt"

func main() {
	// var names [3]string

	// names[0] = "sulthan"
	// names[1] = "rafi"
	// names[2] = "hendriansyah"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	var values = [3]int{
		90,
		85,
		90,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

	var biodatas = [4]string{
		"name",
		"job",
		"country",
		"status",
	}
	fmt.Println(biodatas)

}
