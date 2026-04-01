package main
import "fmt"
import "math"
func main() {
    var v, c1, c2, diferenca1, diferenca2, abs1, abs2 float64
    fmt.Scan(&v)
    fmt.Scan(&c1)
    fmt.Scan(&c2)

    diferenca1 = (v - c1)
    diferenca2 = (v - c2)
    abs1 = math.Abs(diferenca1)
    abs2 = math.Abs(diferenca2)

    if abs1 > abs2 {
    fmt.Println("segundo")
    } else if abs1 < abs2  {
        fmt.Println("primeiro")
    } else {
        fmt.Println("empate")
    }
   
}
