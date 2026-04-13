package main
import "fmt"
func main() {
    var p, a int
    fmt.Scan(&p)

    fmt.Print("[ ")
        for a = 0 ; a <= 9 ; a++ {
            if a == p {
                continue
            }
            fmt.Print(a, " ")
        }
        if p != 10 {
            fmt.Print("ceu")
        }
    fmt.Print("]\n")
    
    }
