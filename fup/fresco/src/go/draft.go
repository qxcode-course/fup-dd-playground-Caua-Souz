package main
import "fmt"
import "bufio"
import "os"
import "strings"
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    frase := scanner.Text()
    junto := strings.Join(frase, )

    fmt.Println(junto)
}