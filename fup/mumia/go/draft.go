package main
import "fmt"
func main() {
    var nome string
    var idade int
    fmt.Scan(&nome)
    fmt.Scan(&idade)

    if idade < 12 {
        fmt.Println("<nome> eh ")
    }
    
}
