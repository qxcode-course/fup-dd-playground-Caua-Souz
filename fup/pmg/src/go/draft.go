package main
import "fmt"
func main() {
    var q int
    var soma, media float64
    fmt.Scan(&q)

    soldados := make([]float64, q)
    for i := 0; i < q; i++ {
        fmt.Scan(&soldados[i])
    }
    for _, valor := range soldados {
        soma+= valor
    }
     media = soma/float64(q)
    fmt.Printf("%.2f",media)

    for i := 0; i < q; i++ {
        if soldados[i] < media {
              soldados[i] = 'P'
        } else if soldados[i] > media {
            soldados[i] = 'G'
            } else {
                soldados[i] = 'G'
            }
    }
    fmt.Print(soldados)

    
}
   
