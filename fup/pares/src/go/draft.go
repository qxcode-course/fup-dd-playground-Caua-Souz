package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	acc := 0
	if n1 > n2 {
		fmt.Println("invalido")
	} else {
		for ; n1 <= n2; n1++ {
			if n1%2 == 0 {
				acc += n1

			}

		}
		fmt.Println(acc)
	}
}
