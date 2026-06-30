package main
import "fmt"
import "bufio"
import "os"
import "strings"
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := scanner.Text()
    textomodificad := strings.Join(strings.Fields(texto), " ")
    fmt.Println(textomodificad)
}