package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)


    lista :=  make([]int, n)

    for i := 0 ; i < len(lista) ; i++ {
        fmt.Scan(&lista[i])
    }
    fmt.Print("[")
    for i := 0 ; i < len(lista) ; i++ {
        if i == len(lista) - 1 {
            fmt.Printf("%d" ,lista[i])
        } else {
        fmt.Printf("%d, ", lista[i])
    }
}
    fmt.Print("]\n")
}