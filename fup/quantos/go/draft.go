package main
import "fmt"
func main() {
    var num1, num2, num3 int
    
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Scan(&num3)

    if num1 == num2 && num2 == num3 { //verificar se todos saão iguais
        fmt.Println("3")
    } else if num1 == num2 || num2 == num3 || num1 == num3 { // verificar se há dois iguais
        fmt.Println("2")
    } else {
        fmt.Println("0") // nenhum igual
        
    }
    
}
