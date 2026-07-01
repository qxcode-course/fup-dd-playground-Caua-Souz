package main
import "fmt"
import "bufio"
import "os"
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := scanner.Text()

    scanner.Scan()
    substring := scanner.Text()
    cont := 0

    for i := 0; i <= len(texto)-len(substring); i++ { //enquanto i for menor ou igual ao (largura do texto) - (largura da substring), i incrementa
        if texto[i:i+len(substring)] == substring { //comparamos o pedaço encontrado com o trecho que buscamos, se encontrado...
            cont++ //contador aumenta
        }
    }

    fmt.Println(cont)
}