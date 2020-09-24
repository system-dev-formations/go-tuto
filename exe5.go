package main

import "fmt"

type mystruct struct {
	a   int
	arr []int
}

type A int

func (a A) Foo() {
	a++
	fmt.Println("foo", a)
}

func main() {
	m := mystruct{}
	//m.arr := make([]int, 5) //compilation error because m.arr is already declared.
	m.arr = make([]int, 5) //compiles
	//--------------
	var a A
	fmt.Println("before", a) // 0
	a.Foo()                  // 1
	fmt.Println("after", a)  // 0
}
