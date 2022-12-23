package main

import (
	"log"
)

func main() {
	// LIS ex1
	nums1 := []int{5, 7, 4, -3, 9, 1, 10, 4, 5, 8, 9, 3}
	expected := 6
	got := lis(nums1)
	if expected != got {
		log.Printf("[FAIL]: Expected %d, got %d", expected, got)
	}
}

func fib1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	F := make([]int, n+1)

	F[0] = 0
	F[1] = 1

	for i := 2; i <= n; i++ {
		F[i] = F[i-1] + F[i-2]
	}

	return F[n]
}

func lis(nums []int) int {
	// Find longest increasing subsequence in nums
	// return length
	maxLen := 0
	T := make([]int, len(nums))
	T[0] = 1 // base case
	T[1] = 2
	T[2] = 1 // this is true because the solution MUST include nums[2], according to our recurrence.

	return maxLen
}
