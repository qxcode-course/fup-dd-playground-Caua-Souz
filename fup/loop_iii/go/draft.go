package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    
fmt.Print("[", " ")
    for n1 := n1 ; n2 < n1; n1--{
    fmt.Print(n1 ," ")
}
    fmt.Println("]")
}