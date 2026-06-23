package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    palavra := scanner.Text() 
    var ind, qtd int
    fmt.Scan(&ind)
    fmt.Scan(&qtd) 

    if ind <= 0 || qtd <= 0 || ind >= len(palavra) {
        fmt.Print("")
        return
    }

    substr := ""
    for i := ind ; i < len(palavra) && i < ind+qtd; i++ {
        substr += string(palavra[i])
    }
    fmt.Print(substr)

}