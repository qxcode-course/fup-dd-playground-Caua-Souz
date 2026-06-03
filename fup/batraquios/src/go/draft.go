    package main
    import "fmt"
    func main() {
        var ne1, ne2 int


        fmt.Scan(&ne1) //escanear a quantidade de elementos do primeiro vetor
        lista1 := make([]int, ne1) //criar o vetor baseado na quantidade escaneada
        for i1 := 0 ; i1 < ne1 ; i1++ { //enquanto i1 for menor que a quantidade de vetores, i1 incrementa
            fmt.Scan(&lista1[i1]) //escaneia o vetor continuamente

        }
        fmt.Scan(&ne2) //mesma coisa só que com outro contador, quantidade e lista
        lista2 := make([]int, ne2)
        for i2 := 0 ; i2 < ne2 ; i2++ { 
        fmt.Scan(&lista2[i2])
    }

        for _, x := range lista1 { //x é igual ao tamanho da primeira lista
            contido := false 

        for _, y := range lista2 { //y é igual ao tamanho da segunda lista
            if x == y { //se x for igual a y
                contido = true //então os elementos da primeira lista estão contidos na segunda
                break //quebra pois não é necessário continuar o processo
            }
        }

        if !contido { //se não está contido, imprimir não
            fmt.Println("nao")
            return //pular 
    }
    }
    fmt.Println("sim")
    }