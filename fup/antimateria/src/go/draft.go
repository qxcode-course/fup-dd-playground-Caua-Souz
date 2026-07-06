package main
import "fmt"
import "bufio"
import "os"
func main() {
    leitor := bufio.NewReader(os.Stdin) //criamos um leitor especifico
    var palv1, palv2 string //palavra a, palavra b
    fmt.Fscanln(leitor, &palv1) //lemos as duas palavras com fscan, que lê peloos endereços da memória
    fmt.Fscan(leitor, &palv2)

    A := []rune(palv1) //convertemos as palavras em vetores de runas
    B := []rune(palv2)

    for len(A) > 0 && len(B) > 0 && A[len(A) - 1] == B[0] { //aqui detectamos se a última letra de A é igual a primeira de B
        A = A[:len(A)-1] //removemos a última letra usando map
        B = B[1:] //pega do segundo elemento arté o final
    }
    result := append(A, B...) //juntamos as palavras
    fmt.Println(string(result)) //imprimindo convertido em string
}