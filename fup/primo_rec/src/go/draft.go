package main
import "fmt"

func primo(n int) bool {
    if n <= 1 { 
        return false //retornar falso se for igual ou menor que um
    }
    for i := 2; i*i <= n; i++ { //o i := 2 significa que o código vai procurar de 2 em 2 até aparecer a raíz de n
        if n%i == 0 { //no caso, se acharmos um divisor...
            return false //retornar falso
        }
    
    }
    return true
} 
func main() {
    var n int
    fmt.Scan(&n)
    if primo(n){
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}