package main

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) > 1 {
		mid := len(buildings) >> 1
		m1 := getSkyline(buildings[:mid])
		m2 := getSkyline(buildings[mid:])
		return merge(m1, m2)
	}
	return [][]int{
		{buildings[0][0], buildings[0][2]},
		{buildings[0][1], 0},
	}
}

// [ [2, 10], [9, 0] ] [ [3, 15], [7, 0] ]
// [ [2, 10], [3, 15], [7, 10], [9, 0] ]
func merge(m1, m2 [][]int) (ret [][]int) {
	i, j, h1, h2 := 0, 0, 0, 0
	for i < len(m1) && j < len(m2) {
		var bd []int
		if m1[i][0] < m2[j][0] {
			h1 = m1[i][1]
			bd = []int{m1[i][0], max(h2, h1)}
			i++
		} else if m1[i][0] > m2[j][0] {
			h2 = m2[j][1]
			bd = []int{m2[j][0], max(h2, h1)}
			j++
		} else {
			h1 = m1[i][1]
			h2 = m2[j][1]
			i++
			j++
			bd = []int{m1[i][0], max(h2, h1)}
		}
		if len(ret) == 0 || ret[len(ret)-1][1] != bd {
			ret = append(ret, bd)
		}
	}
	if i == len(m1) {
		return append(ret, m2[j:]...)
	}
	return append(ret, m1[i:]...)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
