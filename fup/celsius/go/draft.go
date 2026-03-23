package main
import "fmt"
func main() {
    var celsius, fahrenheit float64 //tipo das variáveis
    fmt.Scan(&celsius) //ler o celsius
    fahrenheit = (1.8 * celsius) + 32 //converter celsius para fahrenheit
    fmt.Printf("%.6f\n", fahrenheit) //mostrar o resultado com 6 casas decimais ("%.6f\n")
}
