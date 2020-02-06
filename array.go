package main

import "fmt"

func main() {
	var a [5]int
	b := [3]string{"e", "hi", "hihihi"}

	//slide
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var z [2][3]int

	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	fmt.Println("สมาชิก Array มีทั้งหมด => ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println("สมาชิกใน Array ตัวที่ ", i+1, " = ", a[i])
	}

	for i := range a {
		fmt.Println("for 2 =>", a[i])
	}

	for i, v := range a {
		fmt.Println("for 3=>", i, " -- ", v)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	for i := 0; i < len(z); i++ {
		for j := 0; j < len(z[i]); j++ {
			z[i][j] = j
		}
	}

	fmt.Println("aaaa =>", a)
}
