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
	Y := []string{"A", "B", "E", "C", "B", "A"}

	expected = 4
	got = lcs(X, Y)

	if expected != got {
		log.Printf("[FAIL]: Expected %d, got %d", expected, got)
		os.Exit(1)
	}

	log.Printf("[SUCCESS]: Example 2")

	// Example 3: DPV 6.1
	ex3X := []int{5, 15, -30, 10, -5, 40, 10}
	// ex3Sol := []int{10, -5, 40, 10}
	ex3_got := DPV6_1(ex3X)
	fmt.Println(ex3_got)
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
	L := make([][]int, len(X)+1)
	for i, _ := range L {
		L[i] = make([]int, len(Y)+1)
	}

	i, j := 1, 1
	for i < len(L) {
		for j < len(L[i]) {
			if X[i-1] == Y[j-1] {
				L[i][j] = 1 + L[i-1][j-1]
			} else {
				a := L[i-1][j]
				b := L[i][j-1]

				if a > b {
					L[i][j] = a
				} else {
					L[i][j] = b
				}
			}
			j++
		}
		j = 1
		i++
	}

	printTable(L)

	return L[len(L)-1][len(L[len(L)-1])-1]
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

func DPV6_1(X []int) []int {
	T := make([]map[string]interface{}, len(X))
	for i, _ := range T {
		T[i] = make(map[string]interface{})
		T[i]["subsequence"] = make([]int, 1)
		T[i]["sum"] = 0
	}

	T[0]["subsequence"] = []int{X[0]}
	T[0]["sum"] = X[0]

	i := 1
	for i < len(X) {
		Xi := X[i]

		a := T[i-1]["sum"].(int) + Xi
		b := Xi

		if a > b {
			copy(
				T[i]["subsequence"].([]int),
				T[i-1]["subsequence"].([]int),
			)

			T[i]["subsequence"] = append(T[i]["subsequence"].([]int), Xi)
			T[i]["sum"] = a
		} else {
			T[i]["subsequence"] = []int{Xi}
			T[i]["sum"] = Xi
		}
		i++
	}

	tableMaxIdx := 0
	max := 0
	for i, entry := range T {
		if entry["sum"].(int) > max {
			tableMaxIdx = i
		}
	}

	fmt.Println(T)
	return T[tableMaxIdx]["subsequence"].([]int)
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

func printTable(T [][]int) {
	for _, row := range T {
		fmt.Println(row)
	}
}
