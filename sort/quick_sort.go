package sort

import "fmt"

// 给定一个数组，长度n，获取第k个最大值
func searchMax(arr []int, n, k int) int {
	searchQuickSort(arr, 0, len(arr)-1, k)

	fmt.Println(arr)

	return arr[k-1]
}

func searchQuickSort(arr []int, start, end, k int) {
	if start >= end {
		return
	}

	m := Mid(arr, start, end)

	if m == k-1 {
		return
	} else if m < k-1 {
		quickSort(arr, m+1, end)
	} else {
		quickSort(arr, start, m-1)
	}
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	m := Mid(arr, start, end)

	quickSort(arr, start, m-1)
	quickSort(arr, m+1, end)
}

func Mid(arr []int, start, end int) int {
	v := arr[end]

	var l, r = start, end - 1

	for l <= r {
		for l <= r && arr[l] > v {
			l++
		}

		for l <= r && arr[r] <= v {
			r--
		}

		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	arr[l], arr[end] = arr[end], arr[l]

	return l
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
