package pkg

func FixedPoint(list []int) interface{} {
	for index, item := range list {
		if index == item {
			return index
		}
	}
	return false
}
