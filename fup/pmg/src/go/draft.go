package main
import "fmt"

func main() {
    var q int
    var soma, media float64
    fmt.Scan(&q)

    soldados := make([]float64, q) //criando a lista da altura doas soldados
    classificacao := make([]string, q) //lista para a classificacao das alturas
    for i := 0; i < q; i++ {
        fmt.Scan(&soldados[i]) //escanear a lista das alturas no buffer
    }
    for _, valor := range soldados { //para cada altura na lista
        soma+= valor //somar as alturas
    }
     media = soma/float64(q) //calcular a média das alturas e printar com duas casas decimais
    fmt.Printf("%.2f\n",media)

    for i := 0; i < q; i++ {
        if soldados[i] < media { //se a altura do soldado fpr menor que a média...
              classificacao[i] = "P" //ele recebe p
        } else if soldados[i] > media { //se for maior que a média...
            classificacao[i] = "G" //recebe G
            } else {
                classificacao[i] = "M" //caso contrário recebe M
            }
    }

    for i := 0; i < q; i++ {
        fmt.Print(classificacao[i])
        if i < q-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
}
   
