package main

import (
	"fmt"
	"sort"
)

/*
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}
	output := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			output = append(output, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	output = append(output, []int{start, end})
	return output
}

func main() {
	// inputs := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	inputs := [][]int{{1, 4}, {1, 4}}
	fmt.Println(merge(inputs))
}
