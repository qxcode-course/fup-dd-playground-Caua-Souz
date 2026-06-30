package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    solteiros := make(map[int]int) //criando um mapa de inteiros em um vetor de solteiros
    casais := 0 //casais iniciam em zero

    for i := 0 ; i < n ; i++ {
        var especie int //criar um delimitador para o mapa 
        fmt.Scan(&especie)
        if solteiros[-especie] > 0 { //se existir femêas solteiras daquela espécie...
            solteiros[-especie]-- //decrementamos o valor das solteiras
            casais++ //criamos um casal

        } else { //se não...
            solteiros[especie]++ //os solteiros machos crescem
        }
    }
    fmt.Println(casais)
}