package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    vetor := make([]int, n)

    for i := 0 ; i < n ; i++ {
        fmt.Scan(&vetor[i])
    }

    Jedi := 0
    Sith := 0

    for i := 0 ; i < n ;i++ {
        if i < (n/2) {
            Jedi += vetor[i]
        } else if i >= (n/2) {
            Sith += vetor[i]
        }
    }
    if Jedi > Sith {
    fmt.Println("Jedi")
} else if Sith > Jedi{
    fmt.Println("Sith")
} else if Jedi == Sith{
    fmt.Println("Empate")
}
}