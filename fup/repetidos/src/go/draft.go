
package main
import "fmt"
func main() {
    var p, n int
    fmt.Scan(&p, &n)
    
    lista := make([]int, n) //criar uma lista a partir de n inteiros
    contador := 0
    for i := 0 ; i < n ; i++ {
        fmt.Scan(&lista[i])
        if lista[i] == p { //se um indíce da lista for igual ao número escolhido...
            contador++ //...o contador incrementa
        }
    }
    

    fmt.Println(contador)
}
