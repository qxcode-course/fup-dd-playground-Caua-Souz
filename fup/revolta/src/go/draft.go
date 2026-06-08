package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    somapares := 0
    somaimpares := 0


    lista := make([]int, n)
    for i := 0; i < n ; i++ {
        fmt.Scan(&lista[i])
        if lista[i] % 2 == 0 { //se o indíce atual da lista for par...
            somapares = somapares + lista[i] //o contador dos pares armazena esse número para ser somado
        } else {
            somaimpares = somaimpares + lista[i] //o contador dos ímpares armazena esse número para ser somado
        }
    }
        if somaimpares > somapares {//se a soma dos ímpares for maior que a dos pares...
            fmt.Println("soldados") //printar soldados
        } else if somapares > somaimpares { //se não...
            fmt.Println("rebeldes") //printar rebeldes...
        } else {
            fmt.Println("empate") //se forem iguais
        }
    
}