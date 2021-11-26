func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 1 {
		return len(s)
	}

	chars := make(map[byte]int)
	left := 0
	right := 2
	max := 2
	chars[s[0]] = 0
	chars[s[1]] = 1

	for right < len(s) {
		if _, ok := chars[s[right]]; ok || len(chars) < 2 {
			chars[s[right]] = right
		} else {
			leftmost := len(s)
			var leftCh byte
			for c := range chars {
				if chars[c] < leftmost {
					leftmost = chars[c]
					leftCh = c
				}
			}
			delete(chars, leftCh)
			chars[s[right]] = right
			if (right - left) > max {
				max = right - left
			}
			left = leftmost + 1
		}
		right++
	}

	if max < right-left {
		return right - left
	}

	return max
}