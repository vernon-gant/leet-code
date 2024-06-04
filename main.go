package main

import (
	"algos/modified_binary_search"
    "fmt"
)

func main() {
	obj := modified_binary_search.Search2([]int{1,0,1,1,1}, 0)
	fmt.Print(obj)
}
