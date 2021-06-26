package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	mPool = flag.Int("M", 35, "amount of numbers in the main pool")
	mNum  = flag.Int("m", 7, "amount of numbers to select from the main pool")
	sPool = flag.Int("S", 20, "amount of numbers in the supplementary pool")
	sNum  = flag.Int("s", 1, "amount of numbers to select from the supplementary pool")
	games = flag.Int("g", 4, "number of games to play")
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	for i := 1; i <= *games; i++ {
		w := generate(*mPool, *mNum, *sPool, *sNum)
		for _, n := range w.numbers {
			spaces := " "
			if n < 10 {
				spaces = "  "
			}
			fmt.Printf("%d%s", n, spaces)
		}
		if len(w.supplementaries) > 0 {
			fmt.Print("| ")
			for _, s := range w.supplementaries {
				fmt.Printf("%d ", s)
			}
		}
		fmt.Println()
	}
}
