package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint16
}

func (p *Person) sayHello() {
	fmt.Println("hello", p.Name)
}

func main() {
	var guy = new(Person)
	guy.Name = "david"
	guy.sayHello()

}
