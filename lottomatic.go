package main

import (
	"math/rand"
	"sort"
	"time"
)

// winning holds the winning lotto numbers
type winning struct {
	numbers         []int
	supplementaries []int
}

// generate a random slices of numbers
// poolSize: the larget number in the pool
// main: the number of numbers in the main pool
// supSize: the larget number in the supplementary pool
// sups: the number of supplementary numbers or powerballs
// dupSup: do the supplementary numbers come from the same
// pool or a different pool to the main numbers
// returns a struct of winning numbers
func generate(poolSize int, main int, supSize int, sups int) winning {

	w := winning{}
	w.numbers = getNumbers(poolSize, main)
	if supSize > 0 {
		w.supplementaries = getNumbers(supSize, sups)
	}
	return w
}

// getNumbers
func getNumbers(size int, howMany int) []int {
	low := 1
	numbers := []int{}
	for i := 0; len(numbers) < howMany; i++ {
		n := randomNumber(low, size)
		if !find(numbers, n) {
			numbers = append(numbers, n)
		}
	}
	sort.Ints(numbers)
	return numbers
}

// randomNumber returns a random number
// in the range of low and high numbers
func randomNumber(low int, high int) int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(high+1-low) + low

	return v
}

// find searches for a value: v in
// a slice: s, returning true if found
// and false if not
func find(s []int, v int) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}
