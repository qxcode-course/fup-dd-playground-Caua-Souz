package main
import "fmt"
func main() {
    lista := make([]int, 5) //criar uma lista de cinco inteiros
    for i := 0 ; i < 5 ; i++ { //para i sendo 0, enquanto i for menor que 5, incremente o i
        fmt.Scan(&lista[i]) //ler o índice i da lista
    } 
    menor := lista[0] //identificando o menor valor da lista
        for _, i2 := range lista { // i2 (novo contador) é igual ao tamanho da lista
            if i2 < menor { //se esse contador for menor do que o primeiro índice da lista 
                menor = i2 //igualamos os valores entre as variáveis
            }
        }
    fmt.Println(menor)
}
