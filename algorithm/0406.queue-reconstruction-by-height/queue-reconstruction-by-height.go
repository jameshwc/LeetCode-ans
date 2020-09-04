/*

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

*/

// [[4,4], [5,0], [5,2], [6, 1], [7,0], [7,1]]
//
// [[]]

package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i := range people {
		idx := people[i][1]
		if idx+1 != len(res) {
			copy(res[idx+1:], res[idx:])
		}
		res[idx] = people[i]
	}
	return res
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
