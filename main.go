package main

import (
	"algos/recursion"
	"fmt"
)

func main() {
	result := recursion.GenerateAllParanthesis(3)
	for _,entry := range result {
		fmt.Println(entry)
	}
}
