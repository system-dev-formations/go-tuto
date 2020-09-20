package main

import "fmt"

type Cruncher func(int) int

func mul(n int) int {
	return n * 2
}
func add(n int) int {
	return n + 100
}
func sub(n int) int {
	return n - 1
}
func crunch(nums []int, a ...Cruncher) (rnums []int) {
	// create an identical slice
	rnums = append(rnums, nums...)
	for _, f := range a {
		for i, n := range rnums {
			rnums[i] = f(n)
		}
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("value= %v\n", crunch(nums, mul, add, sub))
}
