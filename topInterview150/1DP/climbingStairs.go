package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var supp []int
	supp = append(supp, 1, 2)
	for i := 2; i <= n; i++ {
		supp = append(supp, supp[i-1]+supp[i-2])
	}
	return supp[n-1]
}
