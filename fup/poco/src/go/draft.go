package main

import "fmt"

func main() {
	var P, S, E int

	fmt.Scan(&P)
	fmt.Scan(&S)
	fmt.Scan(&E)

	pos := 0
	salto := S

	for {
		topo := pos + salto

		if topo >= P {
			fmt.Printf("%d saiu\n", pos)
			break
		}

		fmt.Printf("%d %d\n", pos, topo)

		pos = topo - E

		if pos < 0 {
			fmt.Printf("%d morreu\n", pos)
			break
		}

		salto -= 10
	}
}