package sliding_window

func findRepeatedDnaSequences(s string) []string {
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