package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan()
   frase := scanner.Text()

   
   vogais := make([]rune, 0)
   consoantes := make([]rune, 0)
   
    for _, letra := range frase {
        switch letra {
        case 'a', 'e', 'i', 'o', 'u': //no caso da letra ser as vogais...
            vogais = append(vogais, letra) //acrescentar na lista de vogais...
        case ' ' : //no caso de espaço, pular espaço

        default: //último caso, acrescentar consoantes na lista das consoantes
            consoantes = append(consoantes, letra)
        }

    }
    //converter de runa pra letra antes de imprimir
    fmt.Println(string(vogais))
    fmt.Println(string(consoantes))
}