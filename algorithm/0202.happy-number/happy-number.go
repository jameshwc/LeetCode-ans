func isHappy(n int) bool {
	path := make(map[int]bool)
	for n != 1 {
		if path[n] {
			return false
		}
		path[n] = true
		c := 0
		for n != 0 {
			d := n % 10
			c += d * d
			n = n / 10
		}
		n = c
	}
	return true
}

func isHappyWithoutMap(n int) bool {
	findNext := func(num int) int {
		c := 0
		for num != 0 {
			d := num % 10
			c += d * d
			num = num / 10
		}
		return c
	}

	fast, slow := findNext(findNext(n)), findNext(n)
	for fast != 1 && fast != slow {
		slow = findNext(slow)
		fast = findNext(findNext(fast))
	}
}