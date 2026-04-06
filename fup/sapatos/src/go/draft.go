package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    acc := 0 

    if n1 > n2 { //se n1 for maior que n2, inválido
        fmt.Println("invalido")
    } else {
        for ; n1 <= n2; n1++ { //para cada número entre n1 e n2 incluíndo n2
          if n1%2 == 0 && n1%3 == 0 { //filtrando apenas os divisíveis por 2 e 3
            acc += n1
          }
        }
        fmt.Println(acc)
    }

    

}
