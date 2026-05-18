package main

import (
	"fmt"

)
func main() {
    var cchico, ccebol, animais int
    fmt.Scan(&cchico, &ccebol, &animais)

    lista := make([]string, animais)

    fmt.Println(cchico, ccebol, lista)
}
