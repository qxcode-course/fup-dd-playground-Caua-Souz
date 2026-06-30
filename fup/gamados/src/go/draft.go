package main
import "fmt"
import "bufio"
import "os"
import "strings"
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := strings.Fields(scanner.Text())

    for i := 0; i < len(texto) -1; i++ {
        if texto[i] > texto[i+1] { 
            fmt.Println("nao")
            return
        }
    }
    fmt.Println("sim")
}