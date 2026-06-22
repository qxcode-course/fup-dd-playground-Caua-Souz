package main
import "fmt"


func main() {
    var elem int
    fmt.Scan(&elem)
    vetor := make([]int, elem)
    for i := 0; i < elem; i++ {
        fmt.Scan(&vetor[i])
    }
    movim := 0

    for i := 0; i < elem-1; i++ {
        if vetor[i] < vetor[i+1] && vetor[i+1]-vetor[i] > 1 {
            movim++ 
        } else if vetor[i] > vetor[i+1] && vetor[i]-vetor[i+1] > 1 {
            movim++
        }
    }
     
    fmt.Println(movim)
}