package main
import "fmt"
func main() {
    var caract rune
    fmt.Scanf("%c", &caract)

    if caract >= 'a' && caract <= 'z' {
        caract = caract - ('a' - 'A')
    } else if caract >= 'A' && caract <= 'Z' {
        caract = caract + ('a' - 'A')
    }
    
    fmt.Printf("%c\n", caract)
}