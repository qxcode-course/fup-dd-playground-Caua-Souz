package main
import "fmt"
func main() {
    var num1, num2, soma, subt, multi, divisao int
    var operador string

    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Scan(&operador)

    if operador == "+" {
        soma = (num1 + num2) 
        fmt.Println(soma) 
    } else if operador == "-" {
        subt = (num1 - num2)
        fmt.Println(subt)
    } else if operador == "*" {
        multi = (num1 * num2)
        fmt.Println(multi)        
    } else if operador == "/" {
        divisao = (num1 / num2)
        fmt.Println(divisao)
    } else {
    fmt.Println("Hello, World!")
}
}