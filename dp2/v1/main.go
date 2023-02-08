package main

import "log"

func main() {
	knapsack()
}

func knapsack() {
	build := func(values []int, weights []int, bag int) []int {
		T := make([]int, len(values))

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

		if T[len(T)-1] != t.solution {
			log.Fatalf("[Example %d]: expected %d, got %d", i, t.solution, T[len(T)-1])
		}
	}
}
