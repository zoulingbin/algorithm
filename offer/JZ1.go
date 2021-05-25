package offer

/**
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
  [1,2,8,9],
  [2,4,9,12],
  [4,7,10,13],
  [6,8,11,15]
]
给定 target = 7，返回 true。

给定 target = 3，返回 false。
 */

/**
example:
input: 7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
output:true
 */

func Find( target int ,  array [][]int ) bool {
	height := len(array)
	weight := len(array[0])
	if height == 0 || weight == 0 {
		return false
	}
	last := array[height-1][weight-1]
	if target < array[0][0] || target > last {
		return false
	}
	for i := 0; i < height; i ++ {
		for j := 0; j < weight; j++ {
			if array[i][j] == target{
				return true
			}
		}
	}
	return false
}