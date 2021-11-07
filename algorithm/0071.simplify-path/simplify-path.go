package main

import "strings"

func simplifyPath(path string) string {

	paths := strings.Split(string(path), "/")
	simplePaths := make([]string, len(paths))
	top := 0
	for i := 0; i < len(paths); i++ {
		switch paths[i] {
		case ".", "":
			continue
		case "..":
			if top > 0 {
				top--
			}
		default:
			simplePaths[top] = paths[i]
			top++
		}
	}

	var ans strings.Builder
	ans.WriteString("/")
	for i := 0; i < top; i++ {
		ans.WriteString(simplePaths[i])
		if i+1 != top {
			ans.WriteString("/")
		}
	}
	return ans.String()
}
