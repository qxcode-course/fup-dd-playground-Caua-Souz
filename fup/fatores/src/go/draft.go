package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

    fator := 2
    i := 0

    for num != 1 { //enquanto o número for diferente de um...
        if num % fator == 0 { //se o resto da divisão entre num e o fator for igual a zero...
            num = num / fator //acrescente o resultado da divisão entre os dois no próprio numéro
            i++ //incremente i
        } else { //se não...
            if i > 0 { //se a contagem for maior que 0...
                fmt.Println(fator, i) //imprima o fatorando e o número de repetições
            }
            fator++ //incrementar o fator
            i = 0 //zerar a contagem
        }

    }
        if i > 0 {
            fmt.Println(fator, i)
}
}