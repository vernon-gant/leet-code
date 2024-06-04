package main

import (
	"algos/modified_binary_search"
    "fmt"
)

func main() {
	obj := modified_binary_search.FindClosestElements([]int{1,2,3,4,5},4,3)
	fmt.Print(obj)
}
