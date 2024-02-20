package main

import (
	"algos/sliding_window"
	"fmt"
)

func main() {
	repeated := sliding_window.FindRepeatedSequences("AGAGCTCCAGAGCTTG",6)
	fmt.Print(repeated)
}
