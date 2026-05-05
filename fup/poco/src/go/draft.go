package main
import "fmt"
func main() {
    var prof, salto, escorrg, i2, i3 int
    fmt.Scan(&prof)
    fmt.Scan(&salto)
    fmt.Scan(&escorrg)


    i2 = salto
    for i := 0; i <= prof ; i+= salto {
        i2 = salto + i 
            if i2 >= prof {
                salto = salto - 10
                break                
            }
            fmt.Printf("%d %d\n", i, i2)
            i3 = i
            i -= escorrg
    }
        fmt.Printf("%d saiu\n", i3 + salto - escorrg)
}
