package main

import (
	"fmt"
	"strings"
)
func main() {
    var q int
    fmt.Scan(&q)

    cartas := make([]int, q)
    for i := 0; i < q ; i++ {
        fmt.Scan(&cartas[i])
        if cartas[i] == 1 {
            as := strings(cartas[i])
            as[0] = 'A'
        }
    }
    fmt.Println("Hello, World!")
}