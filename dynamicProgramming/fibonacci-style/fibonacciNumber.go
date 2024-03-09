package main

func fib(n int) int {
	var arr []int
	arr = append(arr, 0, 1)
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}

// without slice
func fib(n int) int {
	if n == 1 {
		return 1
	}
	var res int
	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		res = n1 + n2
		n1, n2 = n2, res
	}
	return res
}
