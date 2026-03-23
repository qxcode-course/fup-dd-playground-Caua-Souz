package main
import "fmt"
import "math"

func main() {
    var lado1, lado2, lado3, sp, area float64
    fmt.Scan(&lado1)
    fmt.Scan(&lado2)
    fmt.Scan(&lado3)
    fmt.Scan(&sp)

    sp = (lado1 + lado2 + lado3)/ 2
    area = math.Sqrt(sp*(sp - lado1)*(sp - lado2)* (sp - lado3)) //cálculo da área 

    fmt.Printf("%.2f\n", area) //mostrar o resultado da área com duas casas decimais
    

}
