package sort

import "fmt"

//冒泡排序
func BubbleSort(arr []int)  {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n -i -1; j++{
			if(arr[j]>arr[j+1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
