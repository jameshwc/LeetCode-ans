package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	edge := make([][]int, numCourses)
	for i := range prerequisites {
		indegree[prerequisites[i][0]]++
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
	}
	var queue []int
	for i := range indegree {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		elem := queue[0]
		queue = queue[1:]
		for _, v := range edge[elem] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	for i := range indegree {
		if indegree[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	if canFinish(2, [][]int{{0, 1}}) && canFinish(2, [][]int{{1, 0}}) {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}
