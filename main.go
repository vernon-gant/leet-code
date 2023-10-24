package main

import (
	"algos/fast_slow"
	"fmt"
)

func main() {
	result := fast_slow.CircularArrayLoop([]int{2,-1,1,-2,-2})
	fmt.Print(result)
}
