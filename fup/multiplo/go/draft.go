package main
import "fmt"
func main() {
    var num int

    fmt.Scan(&num)
    if num % 7 ==0 { //verificar se o número é divisível por 7
    fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
    
}
