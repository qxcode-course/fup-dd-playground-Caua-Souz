            package main
            import "fmt"
            func main() {
                var prof, salto, escorr, i2, i3 int
                fmt.Scan(&prof)
                fmt.Scan(&salto)
                fmt.Scan(&escorr)

                i2 = salto //i2 recebe o valor do salto
                for i := 0; i <= prof; i+= salto { //para i igual a 0, enquanto i for igual ou menor do que a profundidade, i aumenta com o valor de salto
                    i2 = salto + i //i2 recebe o valor de salto + i
                        if i2 >= prof { //se i2 for maior ou igual ao poço, o código para
                            break
                        }

                    fmt.Printf("%d %d\n", i, i2) //os %d servem para modificar a exibição das variáveis, uma do lado da outra (o d é pra inteiro)
                    i3 = i //uma nova variável que recebe o valor de i é essencial, já que i não existe fora da repetição
                    i -= escorr //o escorrego do sapo
                }
                fmt.Printf("%d saiu\n", i3 + salto - escorr) 

        }