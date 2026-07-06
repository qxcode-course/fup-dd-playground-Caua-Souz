package main
import "fmt"

type Restaurante struct { //criando uma struct dos restaurantes
    nome string //nome dos restaurantes
    pontos int //notas 
}

func omelhor(restaurantes []Restaurante) string { //a função recebe a lista das structs e retorna o nome do vencedor
    melhor := restaurantes[0]
    for _, v := range restaurantes[1:] {
        if v.pontos > melhor.pontos || v.pontos == melhor.pontos && v.nome < melhor.nome { //essa linha é a determinante para sempre atualizarmos quem é o melhor restaurante, usando v
            melhor = v //aí quem for o melhor, recebe o valor que o v adquiriu
        }
    }
    return melhor.nome
}

func main() {
    var n int
    fmt.Scan(&n)

    restaurantes := make([]Restaurante, n) //criar um vetor baseado na struct
    for i := 0 ; i < n; i++ {
        fmt.Scan(&restaurantes[i].nome, &restaurantes[i].pontos) //escanear os nomes e pontos das structs dentro do buffer
    }
    fmt.Println(omelhor(restaurantes)) //printar quem foi o melhor dentre eles
}