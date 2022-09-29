package memory

func clamp(max, curr int) int {
	if curr < 0 {
		return 0
	} else if curr > max {
		return max
	}

	return curr
}

func wrap(max, curr int) int {
	if curr < 0 {
		return max - curr
	} else if curr > max {
		return curr % (max + 1)
	}

	return curr
}
