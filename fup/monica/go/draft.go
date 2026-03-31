package main
import "fmt"
func main() {
    var M, A, B, C, R int
    fmt.Scan(&M)
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    fmt.Scan(&R)

    C = M - (A + B) 
   if A > B && A > C || B < A && C < A { //comparação entre as idades para descobrir se A for maior
    fmt.Println(A)
   } else if A < B && C < B { //comparação entre as idades para descobrir se B é maior
    fmt.Println(B)
   } else {
    fmt.Println(C)
   }

} 
