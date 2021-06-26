package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	poolSize = flag.Int("M", 35, "numbers in the main pool of numbers")
	mainN    = flag.Int("m", 7, "number of main numbers")
	supSize  = flag.Int("S", 20, "numbers in the pool of supplementary numbers")
	supN     = flag.Int("s", 1, "number of supplementary numbers")
	games    = flag.Int("g", 4, "number of games to play")
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
		w := generate(*poolSize, *mainN, *supSize, *supN)
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
