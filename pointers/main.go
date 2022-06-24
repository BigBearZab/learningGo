package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{
		name: "dz",
		age:  27,
	}
	fmt.Println(p)
	p.updateAge(28)
	fmt.Println(p)
}

func (p *person) updateAge(i int) {
	p.age = i
}
