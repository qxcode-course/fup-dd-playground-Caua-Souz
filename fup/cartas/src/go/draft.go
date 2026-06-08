package main

import (
	"fmt"
    "strings"
)
func main() {
    var q int
    fmt.Scan(&q)

    cartas := make([]int, q) //vetor das cartas

    for i := 0; i < q ; i++ {
        fmt.Scan(&cartas[i])
    }
    resultado := make([]string, q) //outro vetor somente pro switch e case


    for i, carta := range cartas { //substituição do número pela letra
        switch carta {
        case 1:
            resultado[i] = "A"
        case 11:
            resultado[i] = "J"
        case 12:
            resultado[i] = "Q"
        case 13:
            resultado[i] = "K"
        default: //se o resultado não for os números 1, 11, 12, 13, printar normalmente
            resultado[i] = fmt.Sprint(carta)
        }
    }
    fmt.Printf("[%s]\n",strings.Join(resultado, ", ")) //printar corretamnete com vírgula

    }
