
package main
import "fmt"
func main() {
    var p, n int
    fmt.Scan(&p, &n)
    i := 0
    
    lista := make([]int, n)
    for _, p2 := range lista{
        if p2 == p {
            i++
        }
    }

    fmt.Println(i)
}
