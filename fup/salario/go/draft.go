package main
import "fmt"
func main() {
    var sal_a, aum float64
    fmt.Scan(&sal_a)

    if sal_a <= 1000 {     //condição se o salário for igual ou menor que 1000
        aum = (sal_a * 0.2) + sal_a     //calcular os 20% do salário
        fmt.Printf("%.2f\n", aum)    //imprimir o resultado final com duas casas decimais
    } else if sal_a <= 1500 {    //condição se o salário for igual ou menor que 1500
        aum = (sal_a * 0.15 ) + sal_a    //calcular os 15% do salário
        fmt.Printf("%.2f\n", aum)      //imprimir o resultado final com duas casas decimais
    } else if sal_a <= 2000 {     //condição se o salário for igual ou menor que 2000
        aum = (sal_a * 0.10) + sal_a  //calcular os 10% do salário
      fmt.Printf("%.2f\n", aum)   //imprimir o resultado final com duas casas decimais
    } else if sal_a > 2000 {     //condição se o salário for maior que 2000
        aum = (sal_a * 0.5) + sal_a  //calcular os 5% do salário
        fmt.Printf("%.2f\n", aum)    //imprimir o resultado final com duas casas decimais
    }
}
