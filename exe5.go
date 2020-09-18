package main

type mystruct struct {
	a int
	arr []int
}

func main() {
	m := mystruct{}
	//m.arr := make([]int, 5) //compilation error because m.arr is already declared.
	m.arr = make([]int, 5)  //compiles

}
