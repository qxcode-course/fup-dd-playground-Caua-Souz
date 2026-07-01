package main
import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := scanner.Text()
    v := strings.Fields(texto)
    for i := 0; i < len(v); i++ {
        fmt.Print(v[i], " ")
        if i < len(v)-1 {
        fmt.Print(v[i], " ")
    }
}
    fmt.Print(v[len(v)-1])
    fmt.Println()
}