package pkg

func BinarySearch(min, max, destination int) int {
	mod := destination % 10
	if mod == 2 || mod == 3 || mod == 7 || mod == 8 {
		return -1
	}
	for min <= max {
		mid := (min + max) / 2
		if mid*mid == destination {
			return mid
		} else if mid*mid > destination {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
