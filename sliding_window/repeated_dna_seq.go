package sliding_window

func FindRepeatedDnaSequences(s string) []string {
	var resultSequences []string
	sequenceFrequencies := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		currentSequence := s[i : i+10]
		currentFrequency, ok := sequenceFrequencies[currentSequence];
		if ok && currentFrequency == 1 {
			resultSequences = append(resultSequences,currentSequence)
		}
		if !ok {
			sequenceFrequencies[currentSequence] = 1
			continue
		} 
		sequenceFrequencies[currentSequence] = currentFrequency + 1
	}
	return resultSequences
}
