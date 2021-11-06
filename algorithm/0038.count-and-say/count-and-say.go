func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	cnt := 0
	var s strings.Builder
	for i := 0; i < len(prev); i++ {
		cnt++
		if i+1 < len(prev) && prev[i] == prev[i+1] {
			continue
		}
		s.WriteString(strconv.Itoa(cnt) + string(prev[i]))
		cnt = 0
	}
	return s.String()
}