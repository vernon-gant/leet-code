package two_pointers

import "math"

func MinMovesToMakePalindrome(s string) int {
	return minRec(s, 0)
}

func minRec(s string, counter int) int {
	if len(s) <= 1 {
		return counter
	}
	distances := make(map[rune][3]int)
	for i, val := range s {
		distance, found := distances[val]
		if !found {
			distances[val] = [3]int{min(i, len(s) - i - 1), i, -1}
		} else {
			distance[0] += min(i, len(s) - i - 1)
			distance[2] = i
			distances[val] = distance
		}
	}
	minDistance, toLeft, toRight := math.MaxInt, 0, len(s) - 1
	for _, distance := range distances {
		if minDistance < distance[0] || distance[2] == -1 {
			continue
		}
		minDistance = distance[0]
		toLeft = min(distance[1], distance[2])
		toRight = max(distance[1], distance[2])
	}
	sBytes := []byte(s)
	for toLeft > 0 {
		sBytes[toLeft - 1], sBytes[toLeft] = sBytes[toLeft], sBytes[toLeft - 1]
		toLeft--
		counter++
	}
	for toRight < len(s) - 1 {
		sBytes[toRight], sBytes[toRight + 1] = sBytes[toRight + 1], sBytes[toRight]
		toRight++
		counter++
	}
	return minRec(string(sBytes[1 : len(s) - 1]), counter)
}