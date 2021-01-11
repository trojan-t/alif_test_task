package pkg

func FixedPoint(list []int) (int, bool) {
	for index, item := range list {
		if index == item {
			return index, true
		}
	}
	return -1, false
}
