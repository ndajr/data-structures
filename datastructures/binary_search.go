package datastructures

// BinarySearch takes a sorted array and an item.
// If the item is in the array, the function returns its position, otherwise returns -1.
func BinarySearch(list []int, item int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == item {
			return mid
		} else if list[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
