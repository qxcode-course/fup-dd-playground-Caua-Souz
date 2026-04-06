package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

fmt.Print("[", " ")
    for {
        if n1%2 != 0 {
            fmt.Print(n1, " ")
            n1 = n1 + 1
            continue
        } else if n1 == n2 {
        break
        } else {
        n1 = n1 + 1
        continue
    } 
    }
        fmt.Println("]")
}

