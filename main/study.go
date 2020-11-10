package main

import "fmt"

func main() {
	s := Student1{Person{"kai", 24}, 1, "taiwan"}
	fmt.Println(s)

	s2 := Student1{}

	s2.Person = Person{"dad", 100}

	s2.id = 100
	s2.addr = "dasjdksadsak"
	fmt.Println(s2)

}

type Person struct {
	name string
	age  int
}

type Student1 struct {
	Person
	id   int
	addr string
}
