package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin) //definimos a variável scanner como o novo scanner
    scanner.Scan() //ele escaneia o texto dentro do buffer 
    frase := scanner.Text() //guarda a linha lida dentro da variável frase
    runes := []rune(frase) //converte a frase para um vetor de runa

    for i := len(runes)-1; i >= 0; i-- { //aqui é um for decresente, pois precisamos ordenar a palavra
        fmt.Print(string(runes[i])) //convertemos as runas para string e imprimimos
    }
    fmt.Println()
    

}
