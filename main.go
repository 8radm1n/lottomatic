package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	poolSize = flag.Int("p", 50, "numbers in the pool of numbers")
	mainN    = flag.Int("m", 7, "numbers in the main pool of numbers")
	supN     = flag.Int("s", 1, "numbers in the supplementary pool of numbers")
	dup      = flag.Bool("d", false, "do the supplementary numbers come from the same pool as the main numbers")
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
		w := generate(*poolSize, *mainN, *supN, *dup)
		fmt.Printf("%d: ", i)
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
