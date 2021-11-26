func convertToTitle(columnNumber int) string {

	var sb strings.Builder
	var chars []byte

	for i := 0; i < 26; i++ {
		chars = append(chars, byte(i)+'A')
	}

	for columnNumber > 0 {
		sb.WriteRune(rune(chars[(columnNumber-1)%26]))
		columnNumber = (columnNumber - 1) / 26
	}

	var rsb strings.Builder
	s := sb.String()
	for i := len(s) - 1; i > -1; i-- {
		rsb.WriteByte(s[i])
	}
	return rsb.String()
}
