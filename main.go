package main

import (
	"algos/sliding_window"
	"fmt"
)

func main() {
	result := sliding_window.FindRepeatedDnaSequences("AAAAAAAAAAA")
	fmt.Print(result)
}
