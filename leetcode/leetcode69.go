package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var left, right, mid = 1, x, 0

	for left <= right {
		mid = left + (right-left)>>1

		temp := x / mid
		if temp > mid {
			left = mid + 1
		} else if temp < mid {
			right = mid - 1
		} else {
			return mid
		}
	}

	return right
}
