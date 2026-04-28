package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Print("[ ")
    lista := make([]int, n)
    for i := 0 ; i < len(lista) ; i++ {
        fmt.Scan(&lista[ i ])
    }
    for i := 0 ; i < len(lista) ; i++ {
        fmt.Print(lista[i]," ")
    }
    if len(lista) == 0 {
        fmt.Print("")
    }
    fmt.Println("]")
}
