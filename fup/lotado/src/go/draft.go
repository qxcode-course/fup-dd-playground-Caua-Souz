package main
import "fmt"
func main() {
    var c, m, i int
    fmt.Scan(&c)

    for m = 0 ; m < (c*2) ; {
        fmt.Scan(&i)

        m = m + i
        if m == 0 {
            fmt.Println("vazio")
        } else if m < c && m > 0 {
            fmt.Println("ainda cabe")
        } else if m == c || m < (c*2)  {
            fmt.Println("lotado")
        }
    }
    fmt.Println("hora de partir")
    }
    


