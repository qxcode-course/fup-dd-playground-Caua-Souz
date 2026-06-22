package main
import "fmt"

func criarvetor() []int {
    qtd := 0 //criar e escaneia a quantidade
    fmt.Scan(&qtd)
    lista := make([]int, qtd) //criar a lista
    for i := range lista { //para cada elemento na lista...
        fmt.Scan(&lista[i]) //escaneiar a lista
    }
    return lista //retornarr lista
}

func imprimirlista(lista []int) {
    fmt.Print("[ ")
    for _, elemento := range lista {
        fmt.Print(elemento, " ")
    }
    fmt.Println("]")
}

func temounao(lista []int, valor int) bool {
    for _, elemento := range lista { //verificar os elementos da lista
        if elemento == valor { //caso tenha o elemento na lista...
            return true //retornar "true" (tem elemento igual)
        }
    }
    return false //caso contrário retornar false
}
func main() {
    max := 0
    fmt.Scan(&max)
    figurinhas := criarvetor() //criar o vetor com base na função
    album := make([]int, 0)
    repetidos := make([]int, 0)

    for _, elemento := range figurinhas {
        if !temounao(album, elemento) { //se não tiver o elemento no album...
            album = append(album, elemento) //colocar no album
        } else { //caso já tenha...
            repetidos = append(repetidos, elemento) //colocar nas repetidas
        }
    }

    falta := make([]int,0)
        for i := 1; i <= max ; i++ {
            if !temounao(album, i){ //se não tiver o elemento no album se baseando no contador...
                falta = append(falta, i) //colocar na lista o que falta
            }

        }
    imprimirlista(repetidos)
    imprimirlista(falta)
}