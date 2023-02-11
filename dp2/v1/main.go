package main

import (
	"fmt"
	"log"
)

func main() {
	knapsack()
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
