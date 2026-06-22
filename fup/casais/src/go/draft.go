package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    animais := make([]int, n) //cria a lista de animais
    for i := 0; i < n; i++ {
        fmt.Scan(&animais[i]) //escanear a lista no buffer
    } 
    casais := 0 
    contem := make([]bool, n) //lista para ver se aquele animal já tem um par

    for i := 0; i < n; i++ { 
        if contem[i] { //se esse animal já tem um par...
            continue //pular ele
        }
    for i2 := i + 1; i2 < n; i2++ { //loop para analisar possíveis parceiros
        if contem[i2] { //e esse animal já tem um par...
            continue //pular ele
        }
        if animais[i] == -animais[i2] { //verificando se aquele animal tem uma femêa (negativo)
            casais++ //se sm, forma um casal
            contem[i] = true //e põe o status de que esse animal já tem um parceiro
            contem[i2] = true //mesma coisa com o outro animal
            break
        }
    }
}

    fmt.Println(casais)
}