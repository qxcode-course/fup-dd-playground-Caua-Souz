package main

import (
	"fmt"
	"math"
)
func main() {
	var op string
	var num, res float64
	fmt.Scan(&op)
	fmt.Scan(&num)
	fmt.Scan(&res)

	if op == "r" { //ler se a operação é round
	res = math.Round(num) //usando a função round
	fmt.Println(res)
} else if op == "f" { //ler se a operação é floor
	res = math.Floor(num) //usando a função floor
	fmt.Println(res)
} else {
	res = math.Ceil(num) //ler se a operação é ceil
	fmt.Println(res) //usando a função ceil
}
}
