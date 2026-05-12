package main

import "fmt"

func main() {
	var prof, salto, escorrg, i2, i3 int
	fmt.Scan(&prof)
	fmt.Scan(&salto)
	fmt.Scan(&escorrg)

	i3 = 0

	for i := 0; i <= prof; {
		i2 = salto + i3
		if i2 >= prof {
			fmt.Printf("%d saiu\n", i3)
			break
		}
		if i2 < 0 {
			fmt.Printf("%d morreu\n", i2)
			break
		}

		fmt.Printf("%d %d\n", i3, i2)

		i3 = i2 - escorrg

		if i3 < 0 {
			fmt.Printf("%d morreu\n", i3)
			break
		}

		i = i3
		salto = salto - 10
	}
}
