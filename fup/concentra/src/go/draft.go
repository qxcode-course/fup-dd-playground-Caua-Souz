package main
import "fmt"
func main() {
    var n1, n2, i, n int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    n = n2 - n1;
fmt.Print("[ ")
        for  i = 0; i <= n ; i++ {
            fmt.Print(n1, " ", n2, " ")
            n1 += 1
            n2 -= 1
       
                }
fmt.Print("]\n")

            }
        
         



