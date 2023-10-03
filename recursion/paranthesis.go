package recursion

import (
	"fmt"
	"math"
)

func IsBalancedParanth(expression string) bool {
	balance := 0
	return isBalancedParanth(expression,&balance)
}

func isBalancedParanth(expression string,balance * int) bool {
	if *balance < 0 {
		return false
	}
	if len(expression) == 0 {
		return *balance == 0
	}
	if expression[0] == ')' {
		*balance--
	} else if expression[0] == '(' {
		*balance++
	}
	return isBalancedParanth(expression[1:],balance)
}

func GenerateAllParanthesis(openNumber int) {
	expression := "("
	generateAllParanthesis(expression,openNumber * 2 - 1, -1)
}

func generateAllParanthesis(expression string, limit, balance int)  {
	if math.Abs(float64(balance)) > float64(limit / 2) {
		return
	}
	if limit == 0 {
		fmt.Println(expression)
		return
	}
	limit--
	generateAllParanthesis(expression + ")",limit,balance + 1)
	generateAllParanthesis("(" + expression,limit,balance - 1)
	generateAllParanthesis(expression + "(",limit,balance - 1)
	generateAllParanthesis(")" + expression,limit,balance + 1)
}