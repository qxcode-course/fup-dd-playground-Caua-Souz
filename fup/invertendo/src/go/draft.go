package main
import "fmt"

//primeira vez treinando função
func inverter_vetor(lista []int) []int { //(recebe uma lista contendo inteiros "[]int") retorna uma lista contendo inteiros "[]int"
        for i, i2 := 0 , len(lista) - 1; i < i2 ; i, i2=i+1, i2-1 { //para i e i2 igual a zero, alçance de lista menos um, enquanto i for menor que i2, o i2 recebe o valor de i mais um e em seuida menos um
        lista[i], lista[i2] = lista[i2], lista[i] //invertendo os indíces
}
        return lista //retorna a lista, fora do for
}
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Print("[") //primeiro colchete
    lista := make([]int, n) //cria a lista baseada na quantidade escaneada do buffer
    for i := 0 ; i < n ; i++ { //para i igual a zero, enquanto i for menor que a quantidade, incremente i
        fmt.Scan(&lista[i]) //escanear a lista continuamente
    }

    inverter_vetor(lista) //transformar a lista em invertida, usando a função

    for _, elem := range lista{ //para cada elemento da lista
        fmt.Print(" ", elem) //printar os elementos invertidos nus e crus
    }
        fmt.Print(" ]\n") //segundo colchete e quebra de linha
}
