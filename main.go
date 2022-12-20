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
		log.Fatalf("Expected %d, got %d", expected, got)
	}

	log.Println("foo")
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
	return 0
}
