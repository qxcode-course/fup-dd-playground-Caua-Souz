package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

fmt.Print("[", " ")


    if n1 > n2 { //se n1 for maior que n2
        for n1 := n1; n1 > n2; n1-- { //pegar os números de ordem decrescente entre n1 e n2
            fmt.Print(n1," ")  
    } } else {
        for n1 := n1; n1 < n2; n1++ { //pegar os números na ordem crescente entre n1 e n2
            fmt.Print(n1," ")
        }
    }

    fmt.Println("]")
}
