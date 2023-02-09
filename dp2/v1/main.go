package main

import (
	"fmt"
	"log"
)

func main() {
	knapsack()
}

func knapsack() {
	build := func(values []int, weights []int, B int) [][]int {
		T := make([][]int, B+1)
		for i, _ := range T {
			T[i] = make([]int, len(values)+1)
		}

		for i := 1; i <= len(values); i++ {
			T[i]
		}

		return T
	}

	tt := []struct {
		values   []int
		weights  []int
		bag      int
		solution int
	}{
		{
			values:   []int{15, 10, 8, 1},
			weights:  []int{15, 12, 10, 5},
			bag:      22,
			solution: 18,
		},
	}

	for i, t := range tt {
		T := build(t.values, t.weights, t.bag)

		if T[len(T)-1]["value"] != t.solution {
			fmt.Println("DP Table:")
			for _, row := range T {
				fmt.Println(row)
			}
			log.Fatalf("[Example %d]: expected %d, got %d", i+1, t.solution, T[len(T)-1]["value"])
		}
	}
}
