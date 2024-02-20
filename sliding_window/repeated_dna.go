package sliding_window

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
