package main

import (
	"fmt"
	"math"
)
func main() {
    var cchico, ccebolinha, quant int
    fmt.Scan(&cchico, &ccebolinha, &quant)
    animais := make([]string, quant)
    for i := 0 ; i < quant ; i++ {
        fmt.Scan(&animais[i]) //escaneia o que está dentro da lista
    }
    patas := 0
    for _, valor := range animais {
        if valor == "c" { //se o valor da lista for um cavalo
            patas+= 4 //incrementa 4
        } else if valor == "g" { //se for galinha, incrementa 2
            patas+= 2 
        } else if valor == "v" { //se for vaca, incrementa 4
            patas+= 4
        }
    }

    fmt.Println(patas)
    dif1 := math.Abs(float64(cchico) - float64(patas)) //diferença absoluta 1
    dif2 := math.Abs(float64(ccebolinha) - float64(patas)) //diferença absoluta 2
    if dif1 < dif2 {
        fmt.Println("Chico Bento")
    } else if dif2 < dif1 {
        fmt.Println("Cebolinha")
    } else if dif1 == dif2{
        fmt.Println("empate")
    }

    
}
