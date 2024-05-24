package main

import "fmt"

func main() {
	e1 := new(entity)
	e2 := entity{}

	fmt.Println(e1 == nil)
	fmt.Println("e1 =", e1)
	fmt.Printf("%T\n", e1)
	fmt.Println("e2 =", e2)
	fmt.Printf("%T\n", e2)
}

type entity struct{}
