package algo1376

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adj := make([][]int, n)
	for i := range manager {
		if i == headID {
			continue
		}
		adj[manager[i]] = append(adj[manager[i]], i)
	}

	maxInformTime := 0
	var dfs func(int, int)
	dfs = func(r int, currentInformTime int) {
		for _, employee := range adj[r] {
			dfs(employee, currentInformTime+informTime[employee])
		}
		if currentInformTime > maxInformTime {
			maxInformTime = currentInformTime
		}
	}
	dfs(headID, informTime[headID])
	return maxInformTime
}
