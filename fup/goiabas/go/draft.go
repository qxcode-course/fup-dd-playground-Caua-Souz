package main
import "fmt"
import "math"

func main() {
    var cesta, bananas, goiabas, mangas float64
    var resultadop, resultadof float64
    fmt.Scan(&cesta)
    fmt.Scan(&bananas)
    fmt.Scan(&goiabas)
    fmt.Scan(&mangas)
    fmt.Scan(&resultadop)
    fmt.Scan(&resultadof)

    resultadop = (bananas + goiabas + mangas ) / cesta //somar a quantidade das frutas e dividir com a capacidade da cesta
    resultadof = math.Ceil(resultadop) //arrendondar para cima o resultado
    fmt.Println(resultadof)
    
}
