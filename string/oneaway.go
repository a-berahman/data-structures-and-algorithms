package string

// IsOneAway returns true if given strings are one (or zero) edits away
func IsOneAway(s1 string, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)

	if s1len > s2len {
		return isOneInsertAway(s2, s1)
	}

	if s1len < s2len {
		return isOneInsertAway(s1, s2)
	}

	return isOneReplceAway(s1, s2)
}

// isOneInsertAway returns true if given strings are one insert away
// s1 should be the sortest string always
func isOneInsertAway(s1 string, s2 string) bool {
	indexS1 := 0
	indexS2 := 0

	for indexS1 < len(s1) {
		if s1[indexS1] != s2[indexS2] {
			if indexS1 != indexS2 {
				return false
			}
			indexS2++
		} else {
			indexS1++
			indexS2++
		}
	}
	return true
}

// isOneReplceAway returns true if given strings are one replace away
func isOneReplceAway(s1 string, s2 string) bool {
	indexS1 := 0
	indexS2 := 0
	foundDiff := false
	for indexS1 < len(s1) {
		if s1[indexS1] != s2[indexS2] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
		indexS1++
		indexS2++
	}
	return true
}
