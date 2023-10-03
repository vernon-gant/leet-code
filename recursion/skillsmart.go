package recursion

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// 1.
func MyPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return (1 / x) * MyPow2(x, n+1)
	}
	n--
	return x * MyPow2(x, n)
}

// 2. + Assume n is positive
func SumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	lastDigit := n % 10
	n /= 10
	return lastDigit + SumOfDigits(n)
}

// 3.
func ListLength(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	*numbers = (*numbers)[1:]
	return 1 + ListLength(numbers)
}

// 4. + Assume that string must not be trimmed or fine tuned and that all characters are low
func IsPalindrome(testString string) bool {
	if len(testString) == 1 || len(testString) == 0 {
		return true
	}
	if testString[0] != testString[len(testString)-1] {
		return false
	}
	nextString := testString[1 : len(testString)-1]
	return IsPalindrome(nextString)
}

// 5.
func PrintEvenNumbers(numbers *[]int) {
	if len(*numbers) == 0 {
		return
	}
	firstElement := (*numbers)[0]
	if firstElement % 2 == 0 {
		fmt.Println(firstElement)
	}
	*numbers = (*numbers)[1:]
	PrintEvenNumbers(numbers)
}

// 6.
func PrintEvenIndexNumbers(nums []int) {
	printEvenIndexNumbers(&nums, 0)
}

func printEvenIndexNumbers(numbers * []int, idx int) {
	if idx == len(*numbers) {
		return
	}
	if idx%2 == 0 {
		fmt.Println((*numbers)[idx])
	}
	idx++
	printEvenIndexNumbers(numbers, idx)
}

// 7.
func SecondLargestNumber(numbers []int) int {
	return secondLargestNumber(&numbers, numbers[0], math.MinInt64)
}

func secondLargestNumber(numbers * []int, maxValue, secondMax int) int {
	if len(*numbers) == 0 {
		return secondMax
	}
	firstNumber := (*numbers)[0]
	if firstNumber > maxValue {
		secondMax = maxValue
		maxValue = firstNumber
	} else {
		secondMax = max(firstNumber, secondMax)
	}
	*numbers = (*numbers)[1:]
	return secondLargestNumber(numbers, maxValue, secondMax)
}

func RecursiveFileSearch(dirName string) []os.DirEntry {
	fileNames := new([]os.DirEntry)
	recursiveFileSearch(dirName,fileNames)
	return *fileNames
}

func recursiveFileSearch(dirName string, fileNames * []os.DirEntry) {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		log.Println("Failed to read directory:", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			*fileNames = append(*fileNames,entry)
			continue
		}
		path := filepath.Join(dirName, entry.Name())
		recursiveFileSearch(path,fileNames)
	}
}


