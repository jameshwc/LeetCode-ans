func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		if num&1 == 1 {
			cnt++
		}
		num = num >> 1
	}
	return cnt
}