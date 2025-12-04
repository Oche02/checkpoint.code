package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	var words []string
	word := ""

	for _, r := range input {
		if r != ' ' && r != '\t' { 
			word += string(r)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		return
	}

	for i, w := range words {
		if i > 0 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
		for _, r := range w {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
