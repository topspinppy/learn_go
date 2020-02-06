package main

import "fmt"

func main() {
	fmt.Println("Program Calculate Grade")
	fmt.Print("Input Point : ")
	var point int
	fmt.Scanln(&point)
	fmt.Println("=================================")
	if point >= 80 {
		fmt.Println("ผลการเรียนของคุณคือ : ", "A")
	} else if point >= 70 {
		fmt.Println("ผลการเรียนของคุณคือ : ", "B")
	} else if point >= 60 {
		fmt.Println("ผลการเรียนของคุณคือ : ", "C")
	} else if point >= 50 {
		fmt.Println("ผลการเรียนของคุณคือ : ", "D")
	} else {
		fmt.Println("ผลการเรียนของคุณคือ : ", "F")
	}
}
