package main

import (
	"fmt"
	"sort"
)
func main() {
    var n int
    fmt.Scan(&n)

    notas := make([]float64, n)
    for i := 0 ; i < n ; i++ {
        fmt.Scan(&notas[i])
    }
    sort.Float64s(notas)
    fmt.Println(notas)

    if n % 2.0 == 0.0 {
        

    }
}