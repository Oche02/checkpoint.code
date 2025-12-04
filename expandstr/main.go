package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	w, words := "", []string{}

	for _, r := range s {
		if r > ' ' {
			w += string(r)
		} else if w != "" {
			words = append(words, w)
			w = ""
		}
	}
	if w != "" {
		words = append(words, w)
	}
	if len(words) == 0 {
		return
	}

	for i, word := range words {
		if i > 0 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
		for _, r := range word {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
