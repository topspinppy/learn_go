package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b, c = false, 10, "เรียน golang"
	var gos string = "hesheit"
	var gopher string
	gopher = "hihihi"
	//shorthand
	f, d, e := true, 10, "golang is easy"

	fmt.Println("bb=>", f, d, e)
	fmt.Println("aa=>", a, b, c)
	fmt.Println("Hello", gopher, gos)
	fmt.Println("hi!")
	log.Println("")
	log.Println("Hi!!")
	fmt.Println("Hello World")
	println("Hello Hi!")

	fmt.Print("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if len(fruit) == 0 {
		fmt.Printf("ค่าว่างนะค่ะ")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("Apple")
	case "banana":
		fmt.Println("banana")
	case "lemon":
		fmt.Println("lemon")
	default:
		fmt.Println("kuy")
	}
}
