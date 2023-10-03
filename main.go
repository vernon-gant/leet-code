package main

import (
	"algos/recursion"
	"fmt"
)

func main() {
	result := recursion.RecursiveFileSearch("..//")
	for _, entry := range result {
		fmt.Println(entry.Name())
	}
}
