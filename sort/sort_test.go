package sort

import "testing"

func TestQuick(t *testing.T) {
	var arr = []int{3, 1, 15, 12, 9, 6, 10}

	quickSort(arr, 0, len(arr)-1)

	t.Log(arr)
}

func TestSearchMax(t *testing.T) {
	var arr = []int{3, 1, 15, 12, 9, 6, 10}

	t.Log(searchMax(arr, len(arr), 2))
}

func TestSelectSort(t *testing.T) {
	var arr = []int{3, 1, 15, 12, 9, 6, 10}

	selectSort(arr)

	t.Log(arr)
}

func TestBubbleSort(t *testing.T) {
	var arr = []int{3, 1, 15, 12, 9, 6, 10}

	bubbleSort(arr)

	t.Log(arr)
}
