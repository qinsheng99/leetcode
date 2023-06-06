package sort

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		v := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[v] {
				v = j
			}
		}
		swap(arr, i, v)
	}
}
