package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    altura := make([]int, n)
    for i := 0 ; i < n ; i++ {
        fmt.Scan(&altura[i])
    }
    
    fmt.Println("Hello, World!")
}