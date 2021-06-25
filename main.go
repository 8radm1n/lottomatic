package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type winning struct {
	numbers         []int
	supplementaries []int
}

func generate(numRange int, nums int, sups int) winning {
	first := 1
	last := numRange
	result := winning{}
	numbers := []int{}
	supplementaries := []int{}
	for i := 0; len(numbers) < nums; i++ {
		n := randomNumber(first, last)
		if !find(numbers, n) {
			numbers = append(numbers, n)
		} else {
			n := randomNumber(first, last)
			numbers = append(numbers, n)
		}
	}

	for i := 0; len(supplementaries) < sups; i++ {
		n := randomNumber(first, last)
		if !find(supplementaries, n) {
			supplementaries = append(supplementaries, n)
		} else {
			n := randomNumber(first, last)
			supplementaries = append(supplementaries, n)
		}
	}

	sort.Ints(numbers)
	sort.Ints(supplementaries)
	result.numbers = numbers
	result.supplementaries = supplementaries

	return result

}

func randomNumber(first, last int) int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())

	v := rand.Intn(last+1-first) + first

	return v
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	for i := 1; i <= 10; i++ {
		r := generate(60, 7, 1)
		fmt.Printf("%+v\n", r)
	}
}
