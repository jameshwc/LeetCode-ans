func titleToNumber(columnTitle string) int {
	col := 0
	for i := range columnTitle {
		idx := int(columnTitle[i]-'A') + 1
		col = col*26 + idx
	}
}