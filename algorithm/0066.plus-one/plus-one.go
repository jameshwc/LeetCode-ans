func plusOne(digits []int) []int {
	var ans []int
	carry := false
	for i := len(digits) - 1; i > -1; i-- {
		digits[i]++
		if digits[i] > 9 {
			carry = true
			digits[i] = 0
		} else {
			carry = false
			break
		}
	}
	if carry {
		ans = append(ans, 1)
		ans = append(ans, digits...)
	}
	return digits
}