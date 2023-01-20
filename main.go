package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// LIS ex1
	nums1 := []int{5, 7, 4, -3, 9, 1, 10, 4, 5, 8, 9, 3}
	expected := 6
	got := lis(nums1)
	if expected != got {
		log.Printf("[FAIL]: Expected %d, got %d", expected, got)
		os.Exit(1)
	}

	log.Printf("[SUCCESS]: Example 1")

	// Example 2: LCS
	X := []string{"B", "C", "D", "B", "C", "D", "A"}
	Y := []string{"A", "B", "E", "C", "B", "A", "B"}

	expected = 4
	got = lcs(X, Y)

	if expected != got {
		log.Printf("[FAIL]: Expected %d, got %d", expected, got)
		os.Exit(1)
	}

	log.Printf("[SUCCESS]: Example 2")
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

func lcs(X []string, Y []string) int {
	L := make([][]int, len(X))
	for i, _ := range L {
		L[i] = make([]int, len(Y))
	}

	i, j := 1, 1

	for i < len(L) {
		for j < len(L[i]) {
			if X[i] == Y[j] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				a := L[i][j-1]
				b := L[i-1][j]

				if a > b {
					L[i][j] = a
				} else {
					L[i][j] = b
				}
			}
			j++
		}
		i++
	}

	fmt.Println(L)
	return L[len(X)-1][len(Y)-1]
}

func lis(nums []int) int {
	// Find longest increasing subsequence in nums
	// return length
	maxLen := 0
	T := make([]map[string]int, len(nums))
	T[0] = map[string]int{
		"value":  nums[0],
		"length": 1,
	} // base case
	T[1] = map[string]int{
		"value":  nums[1],
		"length": 1 + T[0]["value"],
	}

	for i, num := range nums {
		if i == 0 || i == 1 {
			continue
		}

		// T[i] = 1 + constrained max
		T[i] = map[string]int{
			"value":  num,
			"length": 1 + constrainedMax(T, i),
		}
	}

	for _, v := range T {
		if v["length"] > maxLen {
			maxLen = v["length"]
		}
	}

	return maxLen
}

func constrainedMax(T []map[string]int, i int) int {
	// find a j such that:
	// - j <= i
	// - aj < ai

	max := 0
	for j, v := range T {
		if j <= i && v["value"] < T[i]["value"] && T[i]["length"] >= max {
			max = T[i]["length"]
		}
	}

	return max
}
