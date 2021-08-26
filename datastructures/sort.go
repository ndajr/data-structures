package datastructures

func BubbleSort(arr []int) []int {
	n := len(arr) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}
