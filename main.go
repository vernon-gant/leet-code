package main

import (
    "algos/intervals"
)

func main() {
    result := intervals.Constructor()
    result.Book(48,50)
    result.Book(0,6)
    result.Book(6,13)
    result.Book(8,13)
}
