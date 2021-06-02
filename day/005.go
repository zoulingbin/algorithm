package day

import "math"

func SearchInsert(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := int(math.Floor(float64((low + high) / 2)))
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		}
	}
	return high + 1
}
