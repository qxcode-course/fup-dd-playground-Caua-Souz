package main
import "fmt"

type Pessoa struct { //criamos a struct pessoa
    nome string //nome
    idade int //idade
    sexo byte //sexo em byte, pois ele está em um único caractere (deu erro quando li na primeira vez, achei que seria isso)
}
func main() {
    var n int
    fmt.Scan(&n)

    pessoas := make([]Pessoa, n) //criamos a lista de pessoas baseada na struct
    for i := 0; i < n; i++ {
        var sexo string
        fmt.Scan(&pessoas[i].nome, &pessoas[i].idade, &sexo)
        pessoas[i].sexo = sexo[0]
    }

    IdMaior := -1
    nome := ""

    for _, v := range pessoas {
        if v.sexo == 'f' && v.idade > IdMaior {
            IdMaior = v.idade
            nome = v.nome
        }
    }
    if IdMaior == -1 { //caso o valor armazenado na idade maior não tenha mudado, significa que o código leu só homens
        fmt.Println("nao tem mulher")
    } else {
    fmt.Println(nome)
    }
}