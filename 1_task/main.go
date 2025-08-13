package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) getAge() int {
	return h.age
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			name: "Sanzhar",
			age:  20,
		},
	}
	fmt.Println(a.getAge())
}
