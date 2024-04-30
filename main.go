package main

import (
	"algos/two_heaps"
	"fmt"
)

func main() {
	result := two_heaps.FindMaximizedCapital(2,0,[]int{1,2,3},[]int{0,9,10})
	fmt.Print(result)
}
