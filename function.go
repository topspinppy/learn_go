package main

import "fmt"

func main() {
	// a, b := 1, 2
	// r := add(a, b)
	// fmt.Println(r)

	a := [5]int{}
	fmt.Println(a)
	mutateArray(a[:])
	fmt.Println(a)
}

// func add(x, y int) int {
// 	p := 1
// 	return x + y + p
// }

func mutateArray(a []int) {
	a[0] = 10
	fmt.Println(a)
}
