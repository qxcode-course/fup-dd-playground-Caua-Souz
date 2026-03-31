package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    for a := a ; a < b; a++ {
    fmt.Println(a)
}
}