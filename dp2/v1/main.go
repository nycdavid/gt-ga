package main

import (
	"fmt"
	"log"
)

func main() {
	// knapsack()
	// cmm()
	dpv617()
}

func dpv617() {
	build := func(coins []int, value int) []bool {
		T := make([]bool, value+1)
		T[0] = true
		T[1] = (func() bool {
			for _, coin := range coins {
				if coin == 1 {
					return true
				}
			}

			return false
		})()

		for b := 2; b <= value; b++ {
			possible := false
			for _, c := range coins {
				if c > b {
					continue
				}
				possible = T[b-c]
				if possible {
					continue
				}
			}

			T[b] = possible
		}

		return T
	}

	tests := []struct {
		name     string
		coins    []int
		value    int
		solution bool
	}{
		{
			name:     "Test 1",
			coins:    []int{5, 10, 25},
			value:    40,
			solution: true,
		},
		{
			name:     "Test 2",
			coins:    []int{1, 5, 10, 25},
			value:    41,
			solution: true,
		},
	}

	for _, tt := range tests {
		fmt.Println(tt.name)
		T := build(tt.coins, tt.value)

		for t, row := range T {
			fmt.Printf("[b: %d]: %v\n", t, row)
		}
	}
}

func cmm() {
	build := func(matrices [][][]int) [][]int {
		rootNodeCost := func(i, l, j int) int {
			// l is boundary
			// left subtree is 1..l
			// right subtree is l+1..j
			mi := matrices[i]
			ml := matrices[l]
			mj := matrices[j]

			return len(mi) * len(ml[0]) * len(mj[0])
		}

		N := len(matrices)
		T := make([][]int, N)
		for j, _ := range T {
			T[j] = make([]int, N)
		}

		for s := 1; s < N; s++ {
			for i := 1; i <= N-1-s; i++ {
				j := i + s
				T[i][j] = 1000000

				for l := i; l <= j-1; l++ {
					cur := rootNodeCost(i, l, j) + T[i][l] + T[l+1][j]

					if cur < T[i][j] {
						T[i][j] = cur
					}
				}
			}
		}

		return T
	}

	tests := []struct {
		matrices [][][]int
	}{
		{
			matrices: [][][]int{
				[][]int{
					[]int{1, 2, 3, 4},
					[]int{5, 6, 7, 8},
					[]int{9, 10, 11, 12},
				},
				[][]int{
					[]int{9, 8},
					[]int{7, 6},
					[]int{5, 4},
					[]int{3, 2},
				},
				[][]int{
					[]int{11, 12, 13},
					[]int{14, 15, 16},
				},
				[][]int{
					[]int{20},
					[]int{21},
					[]int{22},
				},
			},
		},
	}

	for _, tt := range tests {
		T := build(tt.matrices)

		for _, row := range T {
			fmt.Println(row)
		}
	}
}

func knapsack() {
	optim := func(T [][]int, values []int, weights []int, b, i int) int {
		wi := weights[i-1]
		if wi > b {
			return T[b][i-1]
		}

		x := values[i-1] + T[(b - wi)][i-1]
		y := T[b][i-1]

		if x > y {
			return x
		} else {
			return y
		}
	}

	build := func(values []int, weights []int, B int) [][]int {
		T := make([][]int, B+1)
		for i, _ := range T {
			T[i] = make([]int, len(values)+1)
		}

		for b := 1; b <= B; b++ {
			for i := 1; i <= len(values); i++ {
				T[b][i] = optim(T, values[0:i], weights, b, i)
			}
		}

		return T
	}

	tests := []struct {
		values    []int
		weights   []int
		bag       int
		solution  int
		maxWeight int
	}{
		{
			values:   []int{15, 10, 8, 1},
			weights:  []int{15, 12, 10, 5},
			bag:      22,
			solution: 18,
		},
	}

	for i, tt := range tests {
		T := build(tt.values, tt.weights, tt.bag)

		if T[tt.bag][len(tt.values)] != tt.solution {
			fmt.Println("DP Table:")
			for _, row := range T {
				fmt.Println(row)
			}
			log.Fatalf(
				"[Example %d]: expected %d, got %d",
				i+1,
				tt.solution,
				T[tt.bag][len(tt.values)],
			)
		} else {
			log.Printf("[Example %d]: PASSED", i+1)
		}
	}
}
