package main

import "fmt"

func main() {
	var num1, num2, num3 int      //variáveis inteiras
	fmt.Scan(&num1, &num2, &num3) //ler as variáveis no local da variável (&)
	fmt.Println(num1 + num2 + num3)
}
