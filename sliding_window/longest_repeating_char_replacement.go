package sliding_window
/*
func CharacterReplacement(s string, k int) int {
    start, end := 0, 0
	biggestWindow, maxOccurence := 0, 0
	currentWindow := make(map[byte]int)
	for {
		for end < len(s) && end - start - maxOccurence <= k {
			incomingLetter, incomingLetterCounter := s[end], 1
			count, found := currentWindow[incomingLetter]
			if end - start - maxOccurence == k && count != maxOccurence {break}
			if found {incomingLetterCounter = count + 1}
			currentWindow[incomingLetter] = incomingLetterCounter
			maxOccurence = max(maxOccurence, incomingLetterCounter)
			end++
		}
		biggestWindow = max(biggestWindow,end - start)
		if end == len(s) {break}
		count, _ := currentWindow[s[start]]
		currentWindow[s[start]] = count - 1
		maxOccurence = findMaxOccurence(currentWindow)
		start++
	}
	return biggestWindow
}

func findMaxOccurence(currentWindow map[byte]int) int {
	maxOccurence := 0
	for _, occurence := range currentWindow {
		maxOccurence = max(maxOccurence,occurence)
	}
	return maxOccurence
}
*/

type AlphaCount = [26]int

func CharacterReplacement(s string, k int) int {
	counts := AlphaCount{}
	var left, right, maxFreq, maxLength int
	for _, char := range s {
		counts[char - 'A']++

		maxFreq = max(maxFreq, counts[char - 'A'])
		for (right - left + 1) - maxFreq > k  {
			counts[s[left] - 'A']--
			left++
		}

		maxLength = max(maxLength, right - left + 1)
		right++
	}
	return maxLength
}