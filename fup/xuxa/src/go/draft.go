package main
import "fmt"

func inverter(frase string) string {
    runas := []rune(frase)
    for i, j := 0, len(runas)-1 ; i < j; i, j = i+1, j-1 {
        runas[i], runas[j] = runas[j], runas[i]
    }
    return string(runas)
}

func main() {
    var frase rune
    fmt.Scan(&frase)
    invertida := inverter
    fmt.Println(invertida)
}
