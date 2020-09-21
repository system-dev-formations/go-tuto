package main

import (
	"fmt"
)

//
type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++
}
func (p Person) DoesNotGrow() {
	p.Age++
}

// swagger:route GET /main_person tests pointer and value
// return the difference of using a pointer or a value
// responses:
//   200: pointer-value

func main() {
	p := Person{"JY", 10}
	p.Grow()
	fmt.Println(p.Age)
	ptr := &p
	ptr.DoesNotGrow()
	fmt.Println(p.Age)
	fmt.Println(ptr.Age)

}
