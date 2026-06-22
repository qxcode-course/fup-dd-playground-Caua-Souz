    package main
    import "fmt"
    func main() {
        lista := make([]int, 5) //criar uma lista com 5 slots

        for i := 0 ; i < 5 ; i++ {
            fmt.Scan(&lista[i]) //escanear os slots da lista no buffer
    }
        maior := lista[0] //definição do maior e menor
        menor := lista[0]
        for i := 1; i < 5; i++ {
            if lista[i] < menor { //se o índice da lista for menor que o menor...
                menor = lista[i] //ele vira o novo menor
            }
        }

        for i := 1; i < 5; i++ { //mesmo procedimento só que para o maior número
            if lista[i] > maior {
                maior = lista[i]
            }
        }
        

        fmt.Println(maior + menor) //printar a soma do maior e menor
}