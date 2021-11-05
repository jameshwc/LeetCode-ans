func myAtoi(s string) int {

	i := 0
	for i < len(s) {
		if s[i] != ' ' {
			break
		}
		i++
	}

	if len(s) == 0 || i == len(s) {
		return 0
	}

	isNegative := false
	if s[i] == '-' {
		isNegative = true
		i++
	} else if s[i] == '+' {
		i++
	}

	var d strings.Builder

	for i < len(s) {
		if s[i] > '9' || s[i] < '0' {
			break
		}
		d.WriteString(string(s[i]))
		i++
	}
	digits := d.String()
	for j := 0; j < len(digits); j++ {
		if digits[j] != '0' {
			digits = digits[j:]
			break
		}
		if j+1 == len(digits) {
			return 0
		}
	}

	isOverflowed := false

	if len(digits) > 10 {
		isOverflowed = true
	} else if len(digits) == 10 {
		bound := "214748364"
		hasSameBound := true
		for j := 0; j < 9; j++ {
			if digits[j] > bound[j] {
				isOverflowed = true
				break
			} else if digits[j] < bound[j] {
				hasSameBound = false
				break
			}
		}
		if !isOverflowed && hasSameBound {
			if digits[9] > '6' && !isNegative {
				isOverflowed = true
			} else if digits[9] > '7' && isNegative {
				isOverflowed = true
			}
		}
	}
	if isOverflowed {
		if isNegative {
			return -(1 << 31)
		}
		return (1 << 31) - 1
	}

	final := 0
	for j := 0; j < len(digits); j++ {
		final = final*10 + (int(digits[j]) - '0')
	}

	if isNegative {
		return -1 * final
	}
	return final
}