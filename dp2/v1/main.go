package main

import (
	"fmt"
	"log"
)

func main() {
	knapsack()
}

func knapsack() {
	build := func(values []int, weights []int, bag int) []map[string]int {
		T := make([]map[string]int, len(values)+1)
		for i, _ := range T {
			T[i] = make(map[string]int)
		}

		// Base case
		if weights[0] <= bag {
			T[1]["value"] = values[0]
			T[1]["weight"] = T[1]["weight"] + weights[0]
		}

		for i := 2; i <= len(values); i++ {
			vi := values[i-1]
			wi := weights[i-1]

			if wi > bag-T[i-1]["weight"] && wi <= bag {
				a := T[i-1]["weight"]
				b := vi

				if a > b {
					T[i]["value"] = a
					T[i]["weight"] = T[i-1]["weight"]
				} else {
					T[i]["value"] = vi
					T[i]["weight"] = wi
				}
			}
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
