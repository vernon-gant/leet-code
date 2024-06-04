package main

import (
	"algos/modified_binary_search"
    "fmt"
)

func main() {
	obj := modified_binary_search.FindClosestElements([]int{25,41,81,85,103,117,319},3,84)
	fmt.Print(obj)
}
