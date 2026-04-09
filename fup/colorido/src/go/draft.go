package main

import (
	"fmt"
)

func main() {
	var n, i int
	var p string
	fmt.Scan(&n)
	fmt.Scan(&p)

	fmt.Print("[ ")
	for i = 0; i <= 10; i++ {
		if i == n {
			continue
		}
		if i == 10 {
			fmt.Print("ceu ")
		} else {
			fmt.Print(i, p, " ")
		}
		if p == "d" {
			p = "e"
		} else {
			p = "d"
		}
	}
	fmt.Print("]\n")
}
