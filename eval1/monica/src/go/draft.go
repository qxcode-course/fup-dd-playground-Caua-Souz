package main
import "fmt"
func main() {
    var m, a, b, c int
    fmt.Scan(&m)
    fmt.Scan(&a)
    fmt.Scan(&b)

    
    c = m - (a + b) 
    m = (a + b + c)
    
    
        if a > b && a > c || b < a && c < a {
            fmt.Println(a)
        } else if b > a && b > c || c < b && a < b {
            fmt.Println(b)
        } else {
            fmt.Println(c)
        }
    
}
