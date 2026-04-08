package main
import "fmt"
func main() {
    var ni, nf int
    fmt.Scan(&ni)
    fmt.Scan(&nf)

    
    for ; ni <= nf ; ni++ {
        if ni%3 == 0 && ni%5 == 0 { //se o primeiro número for divisível por 3 e  por 5
            fmt.Println("zigzag") //imprimir zigzag
        } else if ni%3 == 0 {  //se o primeiro número for divisível apenas por 3
            fmt.Println("zig") //imprimir zig
            continue
        } else if ni%5 == 0 { //se o primeiro número for divisível apenas por 5
            fmt.Println("zag") //imprimir zag
            continue
        } else if ni%3 != 0 || ni%5 != 0 { //se o primeiro número não for divisível nem por 3 ou nem por 5
            fmt.Println(ni) //imprimir o número inalterado
        }
        }
    } //é importante seguir a ordem desses if e else, para manter a prioridade consistente
    

