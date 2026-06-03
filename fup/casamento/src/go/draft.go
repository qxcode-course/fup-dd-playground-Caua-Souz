    package main
    import "fmt"
    func main() {
        lista := make([]int, 5)

        maior := 0
        for i := 0 ; i < 5 ; i++ {
            fmt.Scan(&lista[i])
            if lista[i] > lista[maior] {
            maior = i
        }
    }

        fmt.Println(maior)
}