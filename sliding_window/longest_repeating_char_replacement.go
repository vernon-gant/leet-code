func characterReplacement(s string, k int) int {
  firstLoop := replaceRound(s,k)
	reverseLoop := replaceRound(Reverse(s),k)
	if firstLoop > reverseLoop {
		return firstLoop
	}
	return reverseLoop
}

func replaceRound(s string, k int) int {
	maxLength := 0
	for i := 0; i < len(s) - k; i++ {
		currentChar := s[i]
		replacementCnt, currentLen := 0, 1
		for replacementCnt < k {
			if s[i] != currentChar {
				replacementCnt++
			}	
			currentLen++
		}
		if currentLen > maxLength {
			maxLength = currentLen
		}
	}
	return maxLength
}

func Reverse(s string) string {
	n := len(s)
	reversed := make([]rune, n)
	for _, letter := range(s) {
		n--
		reversed[n] = letter
	}
	return string(reversed[n:])
}