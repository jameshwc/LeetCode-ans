func partitionLabels(S string) []int {
	lastOccur := make([]int, 26)
	for i := range S {
		lastOccur[S[i]-'a'] = i
	}
	var ans []int
	left := 0
	sum := 0
	for left < len(S) {
		right := lastOccur[S[left]-'a']
		for i := left + 1; i < right; i++ {
			if v := lastOccur[S[i]-'a']; v > right {
				right = v
			}
		}
		ans = append(ans, right+1-sum)
		left = right + 1
		sum = right + 1
	}
	return ans
}