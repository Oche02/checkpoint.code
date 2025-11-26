package main

import "fmt"

func main() {
	first := true
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			if j >= i {
				continue
			}
			for k := 9; k >= 0; k-- {
				if k >= j {
					continue
				}
				if !first {
					fmt.Print(",")
				}
				fmt.Printf("%d%d%d", i, j, k)
				first = false
			}
		}
	}
	fmt.Println()
}
