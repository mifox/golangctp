package main

import "fmt"

type Person struct {
}

func (p *Person) Say() {
	fmt.Println("I'm a person.")
}

// 组合
type Student struct {
	Person
}

func main() {
	var s Student
	s.Say()
}
