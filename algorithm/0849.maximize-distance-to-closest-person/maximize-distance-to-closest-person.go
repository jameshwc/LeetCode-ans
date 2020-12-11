func maxDistToClosest(seats []int) int {
	max := 0
	cnt := 0
	for i := range seats {
		if seats[i] == 0 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	max = (max + 1) / 2 // distance
	cnt = 0
	for i := range seats {
		if seats[i] == 1 {
			break
		}
		cnt++
	}
	if cnt > max {
		max = cnt
	}
	cnt = 0
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			break
		}
		cnt++
	}
	if cnt > max {
		max = cnt
	}
	return max
}