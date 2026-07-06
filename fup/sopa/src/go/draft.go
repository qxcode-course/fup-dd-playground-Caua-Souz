package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fib := make([]int64, n+1) //int64 para suportar o tamanho de caracteres
    fib[1] = 1 //o primeiro índice de fibonacci sempre será 1, mas o código lẽ o primeiro indice como 1 ao invés de fib[0]
    if n >= 2 { //mesma coisa com o segundo número da sequência de fibonacci
        fib[2] = 1
    }

    for i := 3 ; i <= n; i++ { //por causa desse erro devemos começar a partir de i := 3
        fib[i] = fib[i-1] + fib[i-2] //o próximo número da sequência de fibonacci sempre vai ser a soma dos dois anteriores
    }
      fmt.Println(fib[n])
}
