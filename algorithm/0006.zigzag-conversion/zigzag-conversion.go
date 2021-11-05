func convert(s string, numRows int) string {
	worker := make([]strings.Builder, numRows)
	i := 0
loop:
	for i < len(s) {
		for j := 0; j < numRows; j++ {
			worker[j].WriteString(string(s[i]))
			i++
			if i == len(s) {
				break loop
			}
		}
		for j := numRows - 2; j > 0; j-- {
			worker[j].WriteString(string(s[i]))
			i++
			if i == len(s) {
				break loop
			}
		}
	}
	var ans strings.Builder
	for i = 0; i < numRows; i++ {
		ans.WriteString(worker[i].String())
	}
	return ans.String()
}