package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "unicode"
func main() {
    scanner := bufio.NewScanner(strings.NewReader())
    scanner.Scan()
    frase := scanner.Text()

    fmt.Println("Hello, World!")
}