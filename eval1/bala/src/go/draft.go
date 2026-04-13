package main
import "fmt"
import "math"
func main() {
    var x1, x2, y1, y2, sub1, sub2, pot1, pot2, dist, raiz float64
    fmt.Scan(&x1)
    fmt.Scan(&y1)
    fmt.Scan(&x2)
    fmt.Scan(&y2)

    sub1 = (x2 - x1)
    sub2 = (y2 - y1)
    pot1 = math.Pow(sub1, 2)
    pot2 = math.Pow(sub2, 2)
    dist = (pot1 + pot2)
    raiz = math.Sqrt(dist)

    fmt.Printf("%.2f\n", raiz)
}
