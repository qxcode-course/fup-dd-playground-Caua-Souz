package main
import "fmt"
func main() {
    var n1, n2, soma, subt, multi, divi, resto int
    fmt.Scan(n1)
    fmt.Scan(n2)
    
    soma = n1 + n2
    subt = n1 - n2
    multi = n1 * n2
    divi = n1 / n2
    resto = n1 % n2

    fmt.Println(soma)
    fmt.Println(subt)
    fmt.Println(multi)
    fmt.Printf("%.2f", float64(divi))
    fmt.Println(resto)
}
