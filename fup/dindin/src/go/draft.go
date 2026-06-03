 package main
import "fmt"
func main() {
    var qtde int
    var sabor, turno string
    fmt.Scan(&qtde)
    choco := 0
    lim := 0
    man := 0
    tar := 0

    for i := 0; i < qtde ; i++ { //enquanto i for menor que a quantdade, i incrementa
        fmt.Scan(&sabor, &turno) //escaneia as variáveis
        if sabor == "c"{ //se o sabor escaneado for chocolate
            choco ++ //chocolate ganha um "ponto"
        } else {
            lim ++ //se não, limão ganha "ponto"
        }
    
    if turno == "m" { //se o turno for manhã, manhã pontua
        man ++
    } else {
        tar ++ //se não, tarde pontua
    }
}
    if choco > lim {
        fmt.Println("c") //printar c se o chocolate for mais vendido
    } else if lim > choco {
        fmt.Println("l") //printar l caso contrário
    } else {
        fmt.Println("empate")
    }
    if man < tar {
        fmt.Println("m") //printar m se tarde tiver mais vendas
    } else if tar < man {
        fmt.Println("t") //caso contrário, printar tarde
    } else {
        fmt.Println("empate")
    
}
}