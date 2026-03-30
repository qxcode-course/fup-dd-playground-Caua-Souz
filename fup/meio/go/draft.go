package main
import "fmt"
func main() {
    var num1, num2, num3 int
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Scan(&num3)

    if num2 > num1 && num1 > num3 { //verificar se o número 1 é o do meio
        fmt.Println(num1) 
    } else if num1 < num2 && num2 < num3 || num1 > num2 && num2 > num3 { // verificar as condições para o número 2 ser o do meio
        fmt.Println(num2)
    } else {
        fmt.Println(num3)
    }
}