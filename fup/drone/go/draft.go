package main
import "fmt"
func main() {
    var A, B, C, H, L, caixa, janela int
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    fmt.Scan(&H)
    fmt.Scan(&L)
    fmt.Scan(&caixa)
    fmt.Scan(&janela)

    janela = (H*L)

    if A*B < janela || B*C < janela { //verificar se duas dimensões são menor que a janela, se sim, a caixa passa
         fmt.Println("S")
    } else {
        fmt.Println("N")
    }

    
}
