package main
import "fmt"
func main() {
    var n1, n2, resultado, resto int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&resultado)

    resultado = (n1 / n2)
    resto  = n1 % n2
    fmt.Println(resultado)
}
