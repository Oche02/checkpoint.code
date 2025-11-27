package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	words := []string{}
	i := 0

	for i < len(s) {
		for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
			i++
		}
		start := i
		for i < len(s) && s[i] > ' ' {
			i++
		}
		if start < i {
			words = append(words, s[start:i])
		}
	}

	if len(words) == 0 {
		return
	}

	out := words[0]
	for _, w := range words[1:] {
		out += "   " + w
	}

	fmt.Println(out)
}
