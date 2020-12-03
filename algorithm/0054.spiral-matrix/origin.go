const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix[0])
	n := len(matrix)
	isVisited := make([][]bool, n)
	for i := range isVisited {
		isVisited[i] = make([]bool, m)
	}
	i := 0
	si, sj := 0, 0
	var ans []int
	for {
		if si < 0 || sj < 0 || si == n || sj == m || isVisited[si][sj] {
			break
		}
		isVisited[si][sj] = true
		ans = append(ans, matrix[si][sj])
		switch i % 4 {
		case RIGHT:
			if sj+1 >= m || isVisited[si][sj+1] {
				si++
				i++
				continue
			}
			sj++
		case DOWN:
			if si+1 >= n || isVisited[si+1][sj] {
				i++
				sj--
				continue
			}
			si++
		case LEFT:
			if sj-1 < 0 || isVisited[si][sj-1] {
				i++
				si--
				continue
			}
			sj--
		case UP:
			if si-1 < 0 || isVisited[si-1][sj] {
				i++
				sj++
				continue
			}
			si--
		}
	}
	return ans
}