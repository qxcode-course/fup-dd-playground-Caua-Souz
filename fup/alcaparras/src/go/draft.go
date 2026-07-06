package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    leitor := bufio.NewScanner(os.Stdin) //criamos um novo scanner
    leitor.Scan() //ele escaneia o texto no buffer
    frase := leitor.Text() //armazenamos esse texto em frase

    leitor.Scan() //escaneamos a letra em buffer
    letra := leitor.Text() //armazenamos em letra
    cont := 0

    for i := 0; i < len(frase); i++{
        if string(frase[i]) == letra { //se a string de frase for igual a letra escolhida...
            cont++ //o contador é incrementado
        }
    }
    fmt.Println(cont)
}