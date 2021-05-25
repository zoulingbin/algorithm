package leetcode

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	maps := make(map[int]int, length)
	for k, v := range nums {
		if j, ok := maps[target-v]; ok {
			return []int{j, k}
		}
		maps[v] = k
	}
	return nil
}

