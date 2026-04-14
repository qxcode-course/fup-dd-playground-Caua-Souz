package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    lista := make([]int, n)
    for i := 1 ; i < len(lista) ; i++ {
        fmt.Scan(&lista[i])
        fmt.Println(i)
    }
}
 