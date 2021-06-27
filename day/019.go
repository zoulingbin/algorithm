package day

// https://leetcode-cn.com/problems/merge-sorted-array/
func merge(arr1, arr2 []int, m, n int) {
	i, j := m -1, n -1
	for k := m + n -1; k >= 0; k-- {
		if j  < 0 || (i >= 0 && arr1[i] >= arr2[j]) {
			arr1[k] = arr1[i]
			i--
		}else{
			arr1[k] = arr2[j]
			j--
		}
	}
}