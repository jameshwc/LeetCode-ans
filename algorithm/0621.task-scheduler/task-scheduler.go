package main

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	count := make([]int, 26)
	for i := range tasks {
		count[tasks[i]-'A']++
	}
	sort.Ints(count)
	i := 25
	max := count[25]
	for i >= 0 && count[i] == max {
		i--
	}
	ans := (max-1)*(n+1) + (25 - i)
	if len(tasks) > ans {
		return len(tasks)
	}
	return ans
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))                               // A B s A B s A B
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2)) // A B C D A B C D A B C E
}
