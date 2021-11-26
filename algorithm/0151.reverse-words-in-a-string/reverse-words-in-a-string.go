func reverseWords(s string) string {
	var words []string
	var word []byte
	for i := range s {
		if s[i] == ' ' {
			if len(word) > 0 {
				words = append(words, string(word))
				word = make([]byte, 0)
			}
			continue
		}
		word = append(word, s[i])
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}
	var revWords []string
	for i := len(words) - 1; i > -1; i-- {
		revWords = append(revWords, words[i])
	}
	return strings.Join(revWords, " ")
}
