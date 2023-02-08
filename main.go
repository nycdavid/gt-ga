package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	// hotelStops6_2()
	// yuckdonalds()
	// corruptTextDoc()
	cuttingCloth()
}

func hotelStops6_2() {
	penalty := func(x int) int {
		return (200 - x) * (200 - x)
	}

	stops := []int{0, 200, 250, 600, 1500, 2000}

	T := make([]map[string]any, len(stops))

	T[0] = make(map[string]any)
	T[0]["penalty"] = 0
	T[0]["stops"] = []int{}

	for j, _ := range T {
		if j == 0 {
			continue
		}

		T[j] = make(map[string]any)

		i := 0
		min := penalty(stops[len(stops)-1])
		for i < j {
			calc := T[i]["penalty"].(int) + penalty(stops[j]-stops[i])
			if calc < min {
				min = calc
			}
			i++
		}
		T[j]["penalty"] = min
	}

	for _, row := range T {
		fmt.Println(row)
	}
}

func yuckdonalds() {
	fmt.Println("=====================")
	fmt.Println("DPV 6.3: Yuckdonald's")
	fmt.Println("=====================")

	K := 4
	m := []int{0, 4, 6, 11, 13, 18}
	p := []int{0, 2, 2, 1, 5, 7}

	alpha := func(m_i int, m_j int) int {
		if math.Abs(float64(m_i)-float64(m_j)) < float64(K) {
			return 0
		} else {
			return 1
		}
	}

	// idx 0-5, 6 elements
	T := make([]int, len(m))
	T[0] = 0
	T[1] = p[1]

	i := 2
	for i < len(m) {
		j := 0
		a := 0

		for j < i {
			if T[j]+alpha(m[j], m[i]) > j {
				j = T[j]
			}

			j++
		}

		a = j
		b := p[i]

		if a >= b {
			T[i] = a
		} else {
			T[i] = b
		}

		i++
	}

	for _, row := range T {
		fmt.Println(row)
	}
}

func corruptTextDoc() {
	dict := func(w string, dictionary []string) bool {
		for _, entry := range dictionary {
			if w == entry {
				return true
			}
		}

		return false
	}

	build := func(input string, dictionary []string) []map[string]any {
		T := make([]map[string]any, len(input)+1)
		T[0] = make(map[string]any)
		T[0]["valid"] = true
		T[0]["words"] = []string{}

		chars := []rune(input)
		for i, _ := range chars {
			// 0th char but T[1]
			exists := false

			T[i+1] = make(map[string]any)
			T[i+1]["words"] = []string{}

			for j := 0; j <= i; j++ {
				if dict(string(chars[j:(i+1)]), dictionary) && T[j]["valid"].(bool) {
					T[i+1]["words"] = append(
						T[j]["words"].([]string),
						string(chars[j:(i+1)]),
					)
					exists = true
				}
			}

			T[i+1]["valid"] = exists
		}

		return T
	}

	tt := []struct {
		input         string
		expectedValid bool
		expectedWords string
		dictionary    []string
	}{
		{
			input:         "itwasthebestoftimes",
			expectedValid: true,
			expectedWords: "it was the best of times",
			dictionary:    []string{"times", "best", "of", "the", "it", "was", "wash"},
		},
		{
			input:         "itwashthebestoftimes",
			expectedValid: true,
			expectedWords: "it wash the best of times",
			dictionary:    []string{"times", "best", "of", "the", "it", "was", "wash"},
		},
		{
			input:         "itwaszbestoftimes",
			expectedValid: false,
			expectedWords: "",
			dictionary:    []string{"times", "best", "of", "the", "it", "was", "wash"},
		},
	}

	for i, t := range tt {
		T := build(t.input, t.dictionary)

		last := T[len(T)-1]

		var expected any
		var got any

		expected = t.expectedValid
		got = last["valid"].(bool)
		if expected != got {
			fmt.Println("DP table:")
			for _, row := range T {
				fmt.Println(row)
			}
			log.Fatalf("[Test %d]: Expected last[\"valid\"] to be %v, got %v", i+1, expected, got)
		}

		expected = t.expectedWords
		got = strings.Join(last["words"].([]string), " ")
		if expected != got {
			fmt.Println("DP table:")
			for _, row := range T {
				fmt.Println(row)
			}
			log.Fatalf("[Test %d]: Expected last[\"words\"] to be %s, got %s", i+1, expected, got)
		}
	}
}

func cuttingCloth() {
	p := [][]int{
		[]int{2, 3},
		[]int{4, 7},
		[]int{1, 4},
	}

	c := []int{2, 4, 6}

	cloth := []int{10, 15}

	build := func(cloth []int, c []int, p [][]int) [][]int {
		T := make([][]int, cloth[1])
		for i, _ := range T {
			T[i] = make([]int, cloth[0])
		}

		return T
	}

	T := build(cloth, c, p)

	for _, row := range T {
		fmt.Println(row)
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

	sol := L[len(L)-1][len(L[len(L)-1])-1]

	return sol
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
			T[i]["subsequence"] = make([]int, len(T[i-1]["subsequence"].([]int)))
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

func matchArray(A []int, B []int) bool {
	if len(A) != len(B) {
		return false
	}

	matches := true

	for i, val := range B {
		if A[i] != val {
			matches = false
			break
		}
	}

	return matches
}
