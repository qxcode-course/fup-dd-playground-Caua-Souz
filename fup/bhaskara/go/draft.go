package main

import (
	"fmt"
	"math"
)

func main() {
    var a, b, c, delta, resultado1, resultado2 float64
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    
   


    delta = (math.Pow(b, 2)) - 4*a*c //equação para achar o delta

    if delta > 0 {
        resultado1 = ((b*-1) + (math.Sqrt(delta)))/ (2*a) //bhaskara da primeira raíz
        resultado2 = ((b*-1) - (math.Sqrt(delta)))/ (2*a) //bhaskara da segunda raíz
        
        fmt.Printf("%.2f\n", resultado1) //printar as raízes com duas casas decimais
        fmt.Printf("%.2f\n", resultado2)
    } else if delta == 0 {
        resultado1 = ((b*-1) + (math.Sqrt(delta)))/ (2*a) //bhaskara da única raíz
        fmt.Printf("%.2f\n", resultado1)
    } else {
        fmt.Println("nao ha raiz real")
    }
}