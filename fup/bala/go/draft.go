package main
import "fmt"
import "math"

func main() {
    var x1, y1, x2, y2, sub1, sub2, po1, po2 ,distancia float64
    fmt.Scan(&x1)
    fmt.Scan(&y1)
    fmt.Scan(&x2)
    fmt.Scan(&y2)

    sub1 = (x2 - x1) //subtração das coordenadas do primeiro tiro
    sub2 = (y2 - y1) //subtração das coordenadas do segundo tiro
    po1 = math.Pow(sub1, 2) //elevar a subtração por dois
    po2 = math.Pow(sub2, 2) //elevar a subtração por dois
    distancia = math.Sqrt(po1 + po2) //raíz quadrada da soma das duas potências

    fmt.Printf("%.2f\n", distancia) //mostrar e formatar a distância
}
