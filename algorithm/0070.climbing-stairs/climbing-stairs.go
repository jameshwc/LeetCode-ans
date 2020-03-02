package main

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	return f[n]
}

