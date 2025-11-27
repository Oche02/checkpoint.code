package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	s := os.Args[1]
	out := ""
	i := 0

	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	started := false
	for i < len(s) {
		if s[i] == ' ' || s[i] == '\t' {
			for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
				i++
			}
			if started && i < len(s) {
				out += " "
			}
		} else {
			out += string(s[i])
			started = true
			i++
		}
	}

	fmt.Println(out)
}
