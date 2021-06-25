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
// sups: the number of supplementary numbers or powerballs
// dupSup: do the supplementary numbers come from the same
// pool or a different pool to the main numbers
// returns a struct of winning numbers
func generate(poolSize int, main int, sups int, dupSup bool) winning {
	low := 1
	high := poolSize
	w := winning{}

	numbers := []int{}
	for i := 0; len(numbers) < main; i++ {
		n := randomNumber(low, high)
		if !find(numbers, n) {
			numbers = append(numbers, n)
		}
	}
	sort.Ints(numbers)
	w.numbers = numbers

	supplementaries := []int{}
	for i := 0; len(supplementaries) < sups; i++ {
		n := randomNumber(low, high)
		if dupSup {
			// Supplementary numbers are drawn from a
			// separate pool to main numbers
			if !find(supplementaries, n) {
				supplementaries = append(supplementaries, n)
			}
		} else {
			// Supplementary numbers are drawn from the
			// same pool as the main numbers
			if !find(numbers, n) && !find(supplementaries, n) {
				supplementaries = append(supplementaries, n)
			}
		}
	}
	sort.Ints(supplementaries)
	w.supplementaries = supplementaries

	return w
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
