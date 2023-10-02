package two_pointers

import "unicode"

func isPalindrome(s string) bool {
	isAlpha := func(sign byte) bool {
		return unicode.IsLetter(rune(sign)) || unicode.IsDigit(rune(sign))
	}
	start, end := 0, len(s) - 1
	for start < end {
		if !isAlpha(s[start]) {
			start++
		} else if !isAlpha(s[end]) {
			end--
		} else {
			if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
				return false
			}
			start++
			end--
		}
	}
	return true
}