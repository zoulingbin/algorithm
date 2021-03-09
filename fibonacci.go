package algorithm

//go实现斐波那契数列的解法
//f[n]=f[n-1]+f[n-2]

//方法1：递归
func fib(n int) int{
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//方法2
func fib2(n int) int{
	cur, next := 0, 1
	for i := 0; i < n; i++ {
		cur, next = next, cur + next
	}
	return cur
}



