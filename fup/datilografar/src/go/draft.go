package main
import "fmt"
import "strings"
func main() {
    var tecla string
    var n string
    fmt.Scan(&tecla)
    fmt.Scan(&n)

    var res strings.Builder

    for _, c := range n {
        if string(c) != tecla {
            res.WriteRune(c)

        }
    }
    resw := strings.TrimLeft(res.String(), "0")

    if resw == ""{
        fmt.Println("0")
    } else {
        fmt.Println(resw)
}
}