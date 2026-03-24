package main
import "fmt"
func main() {
    var n1, n2, nf, media, mediafinal float64
  
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    media = (n1 + n2)/2
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else if 4 > media && media < 7 {
        fmt.Scan(&nf)
        mediafinal = (media + nf)/2
        if mediafinal >= 5 {
            fmt.Println("aprovado na final")
        } 
        } else {
        fmt.Println("reprovado na final")
    }
}
