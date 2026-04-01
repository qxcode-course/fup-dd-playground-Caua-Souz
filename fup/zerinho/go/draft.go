package main
import "fmt"
func main() {
    var d1, d2, d3 int
    fmt.Scan(&d1)
    fmt.Scan(&d2)
    fmt.Scan(&d3)

    if d1 == d2 && d2 == d3 {
        fmt.Println("empate")
    } else if d1 != d2 && d1 != d3 {
        fmt.Println("jog1")
    }  else if d2 != d1 && d2 != d3 {
        fmt.Println("jog2")
    } else {
    fmt.Println("jog3")
}
}