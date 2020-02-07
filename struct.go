package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func mutatePerson(p *person) {
	p.Name = "Hacker"
	fmt.Println("inside mutate", p)
}

type cat struct{}

func (cat) Talk() {
	fmt.Println("Nyaa nyaa")
}

type dog struct{}

func (*dog) Talk() {
	fmt.Println("Wan wan")
}

type talkable interface {
	Talk
}

func main() {
	p1 := person{
		Name:     "TT",
		NickName: "Hi",
	}

	mutatePerson(&p1)
	fmt.Println(p1)
}
