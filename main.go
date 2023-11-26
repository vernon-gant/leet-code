package main

import (
	"algos/sliding_window"
	"fmt"
)

func main() {
	replaced := sliding_window.CharacterReplacement("DIPPITYDIPPPPPPPPPP",4)
	fmt.Print(replaced)
}
