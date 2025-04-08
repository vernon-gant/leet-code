package sliding_window

func CharacterReplacement(s string, k int) int {
	if len(s)-k <= 1 {
		return len(s)
	}

	uniqueLetterIdx := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := uniqueLetterIdx[s[i]]; ok {
			continue
		}
		uniqueLetterIdx[s[i]] = i
	}

	var result int

	for letter, idx := range uniqueLetterIdx {
		result = max(result, processLetter(letter, idx, s, k))
	}

	return result
}

func processLetter(letter byte, idx int, s string, k int) int {
	var i, j, replaced, count, localMax int = idx, idx, 0, 1, 1
	for j+1 != len(s) {
		j++
		if s[j] == letter {
			count++
		} else if replaced <= k {
			replaced++
		} else {
			for count+k != i-j+1 {
				if s[i] == letter {
					count--
				}
				i++
			}
		}
		localMax = max(localMax, j-i+1)
	}
	return localMax
}
