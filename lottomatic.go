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
// mPool: amount of numbers in the main pool
// mNum: amount of numbers to select from the main pool
// sPool: amount of numbers in the supplementary pool
// sNum: amount of numbers to select from the supplementary pool
// returns a struct of winning numbers
func generate(mPool int, mNum int, sPool int, sNum int) winning {

	w := winning{}
	w.numbers = getNumbers(mPool, mNum)
	if sPool > 0 {
		w.supplementaries = getNumbers(sPool, sNum)
	}
	return w
}

// getNumbers returns a random sorted slice of ints
// from a range of 1 to high. howMany is used to
// determine the number of elements in the slice.
func getNumbers(high int, howMany int) []int {
	low := 1
	numbers := []int{}
	for i := 0; len(numbers) < howMany; i++ {
		n := randomNumber(low, high)
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
