package main

import "fmt"

func countingSort(str string) string {
	var bucket [26]int
	for i := range str {
		bucket[str[i]-'a']++
	}
	newStr := []byte(str)
	idx := 0
	for i := 0; i < 26; i++ {
		for bucket[i] != 0 {
			newStr[idx] = uint8(i + 'a')
			idx++
			bucket[i]--
		}
	}
	return string(newStr)
}
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for i := range strs {
		sorted := countingSort(strs[i])
		strMap[sorted] = append(strMap[sorted], strs[i])
	}
	ans := make([][]string, len(strMap))
	id := 0
	for k := range strMap {
		ans[id] = strMap[k]
		id++
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
