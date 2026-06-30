package main
import "fmt"
import "bufio"
import "os"
import "unicode"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := []rune(scanner.Text())

    for i, c := range texto {
        if unicode.IsUpper(c){ //verifica se a letra é maiuscula
            texto[i] = unicode.ToLower(c) //se sim, transformar em miniscula
        } else if unicode.IsLower(c) { //mesma coisa caso seja minuscula
            texto[i] =unicode.ToUpper(c)
        }
    }
        fmt.Println(string(texto))
}