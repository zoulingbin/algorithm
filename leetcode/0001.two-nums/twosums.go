
/**
https://leetcode-cn.com/problems/two-sum/
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和 为目标值target的那两个整数，并返回它们的数组下标
input:[2,7,9,10,16] target:9
output:[0, 1]

input:[3,2,4] target:6
output:[1, 2]
 */
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

