package recursion

import (
	"fmt"
	"math"
)

// 1.
func MyPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return (1 / x) * MyPow2(x, n + 1)
	}
	return x * MyPow2(x, n - 1)
}

// 2. + Assume n is positive
func SumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n % 10 + SumOfDigits(n / 10)
}

// 3.
func ListLength(numbers * []int) int {
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
	if testString[0] == testString[len(testString) - 1] {
		return IsPalindrome(testString[1:len(testString) - 1])
	}
	return false
}


// 5.
func PrintEvenNumbers(numbers []int) {
	if len(numbers) == 0 {
		return
	}
	if numbers[0] % 2 == 0 {
		fmt.Println(numbers[0])	
	}
	PrintEvenNumbers(numbers[1:])
}

// 6.
func PrintEvenIndexNumbers(nums []int) {
	printEvenIndexNumbers(nums,0)
}

func printEvenIndexNumbers(numbers []int, idx int) {
	if idx == len(numbers) {
		return
	}
	if idx % 2 == 0 {
		fmt.Println(numbers[idx])	
	}
	printEvenIndexNumbers(numbers,idx + 1)
}

// 7.
func SecondLargestNumber(numbers []int) int {
	return secondLargestNumber(numbers,numbers[0],math.MinInt64)
}

func secondLargestNumber(numbers []int, maxValue, secondMax int) int {
	if len(numbers) == 0 {
		return secondMax
	}
	firstNumber := numbers[0]
	if firstNumber > maxValue {
		secondMax = maxValue
		maxValue = firstNumber
	} else {
		secondMax = max(firstNumber,secondMax)
	}
	return secondLargestNumber(numbers[1:],maxValue,secondMax)
}

