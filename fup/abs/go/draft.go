package main
import "fmt"
import "math"
func main() {
    var num1, num2, resultado, resultabs float64
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    resultado = (num1 - num2) //subtração dos dois números
    resultabs = math.Abs(resultado) //pegar o absoluto da diferença
    fmt.Println(resultabs) //imprimir o resultado absoluto
    
}
