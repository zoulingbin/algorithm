package sort

//插入排序

func InsertSort(arr []int){
	n := len(arr)
	if n <= 1{
		return
	}
	for i := 1; i <n; i++{
		val := arr[i]
		j := i - 1
		for ; j >=0; j--{
			if arr[j] > val{
				arr[j+1] = arr[j]
			}else{
				break;
			}
		}
		arr[j+1] = val
	}
}
