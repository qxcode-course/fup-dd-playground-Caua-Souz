package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Vogal(r rune) bool {
    r = unicode.ToLower(r)
        return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := []rune(scanner.Text())

    for _, c := range texto {
        if c == ' ' {
            fmt.Print(" ")
        } else if Vogal(c) {
            fmt.Print("v")
        } else {
            fmt.Print("c")
        }
  
    }
    fmt.Println()
}