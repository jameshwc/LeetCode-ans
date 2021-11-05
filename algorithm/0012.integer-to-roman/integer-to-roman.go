func intToRoman(num int) string {
	switch {
	case num >= 1000:
		return "M" + intToRoman(num-1000)
	case num >= 900:
		return "CM" + intToRoman(num-900)
	case num >= 500:
		return "D" + intToRoman(num-500)
	case num >= 400:
		return "CD" + intToRoman(num-400)
	case num >= 100:
		return "C" + intToRoman(num-100)
	case num >= 90:
		return "XC" + intToRoman(num-90)
	case num >= 50:
		return "L" + intToRoman(num-50)
	case num >= 40:
		return "XL" + intToRoman(num-40)
	case num >= 10:
		return "X" + intToRoman(num-10)
	case num >= 9:
		return "IX" + intToRoman(num-9)
	case num >= 5:
		return "V" + intToRoman(num-5)
	case num >= 4:
		return "IV" + intToRoman(num-4)
	case num >= 1:
		return "I" + intToRoman(num-1)
	default:
		return ""
	}
}

func intToRomanGreedy(num int) string {
	words := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i := 0
	var w strings.Builder
	for i := 0; i < len(nums); i++ {
		for num >= nums[i] {
			w.WriteString(words[i])
			num -= nums[i]
		}
	}
	return w.String()
}