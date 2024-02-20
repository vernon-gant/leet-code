package sliding_window

// Leet code
func FindRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	var repeatedSequences []string
	sequencesWithCount := make(map[string]int)
	for i := 0; i <= len(s) - 10; i++ {
		currentSequence := s[i : i + 10]
		count, found := sequencesWithCount[currentSequence]
		if found && count == 1 {
			repeatedSequences = append(repeatedSequences, currentSequence)
		}
		if found {
			sequencesWithCount[currentSequence] = count + 1
			continue
		}
		sequencesWithCount[currentSequence] = 1
	}
	return repeatedSequences
}

// Educative
func FindRepeatedSequences(s string, k int) Set {
	cache, output := *NewSet(), *NewSet()
	windowStart, windowEnd := 0, k

	for ; windowEnd <= len(s); windowStart, windowEnd = windowStart+1, windowEnd+1 {
		currentSubsting := s[windowStart:windowEnd]
		if cache.Exists(currentSubsting) {
			output.Add(currentSubsting)
		}
		cache.Add(currentSubsting)
	}

	return output
}
