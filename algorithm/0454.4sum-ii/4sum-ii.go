func fourSumCount(A []int, B []int, C []int, D []int) int {
	abSum := make(map[int]int)
	cdSum := make(map[int]int)
	for i := range A {
		for j := range B {
			if v, ok := abSum[A[i]+B[j]]; ok {
				abSum[A[i]+B[j]] = v + 1
			} else {
				abSum[A[i]+B[j]] = 1
			}
		}
	}
	for i := range C {
		for j := range D {
			if v, ok := cdSum[C[i]+D[j]]; ok {
				cdSum[C[i]+D[j]] = v + 1
			} else {
				cdSum[C[i]+D[j]] = 1
			}
		}
	}

}