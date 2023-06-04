package main

// 1 2 6 7 8  5
func searchInsert(nums []int, target int) int {
	var l, r = 0, len(nums) - 1

	for l <= r {
		mid := l + ((r - l) / 2)

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return r + 1
}
