package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	from := os.Args[2]
	to := os.Args[3]

	if len(from) != 1 || len(to) != 1 {
		fmt.Println(str)
		return
	}
	rFrom := from[0]
	rTo := to[0]
	found := false
	for i := 0; i < len(str); i++ {
		if str[i] == rFrom {
			found = true
			break
		}
	}
	if !found {
		fmt.Println(str)
		return
	}
	result := []byte(str)
	for i := 0; i < len(result); i++ {
		if result[i] == rFrom {
			result[i] = rTo
		}
	}
	fmt.Println(string(result))
}
