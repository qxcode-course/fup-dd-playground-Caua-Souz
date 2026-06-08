    package main

    import (
        "fmt"
        "sort"
    )
    func main() {
        var n int
        fmt.Scan(&n)

        lista := make([]int, n)
        for i := 0 ; i < n ; i++ {
            fmt.Scan(&lista[i])
        }
        sort.Ints(lista)
        
        for i := 0; i < n; i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(lista[i])
    }
    fmt.Println()

            }
        
