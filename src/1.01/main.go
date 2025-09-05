package main

import "fmt"

type Speaker interface {
	speak() string
}

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h *Human) speak() string {
	return "Hi"
}

func (h *Human) setName(name string) {
	h.name = name
}

func (A *Action) speak() string {
	return "Hi by Actoon"
}

func main() {
	a := Action{
		Human: Human{name: "John"},
	}
	a.setName("Mike")

	var s Speaker = &a

	fmt.Println(a.speak())
	fmt.Println(s)
}
