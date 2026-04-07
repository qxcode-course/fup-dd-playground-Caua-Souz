package main
import "fmt"
func main() {
    var npedra, i int
    fmt.Scan(&npedra)
    i = 0
fmt.Print("[", " ") 
    for ; i <= 9 ; i++ { 
        if i == npedra { //vendo se o contador é igual ao número da pedra.
            continue //se for igual, esse número não entra na contagem, logo o continue é necessário para o "skip".
            }
        fmt.Print(i, " ") //printar a repetição com um espaço à direita.
    }  
   if npedra != 10 {
    fmt.Print("ceu ") //printar "céu" e fechar de forma devida.
   }
   fmt.Println("]")
}
