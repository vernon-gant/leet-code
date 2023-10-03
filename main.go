package main

import (
	"algos/recursion"
	"fmt"
	"sort"
)

func main() {
	fileNames := recursion.RecursiveFileSearch("..//")
	sort.Slice(fileNames, func(i, j int) bool {
		return fileNames[i].Name() < fileNames[j].Name()
	})
	for _, entry := range fileNames {
		fmt.Println(entry.Name())
	}
}
