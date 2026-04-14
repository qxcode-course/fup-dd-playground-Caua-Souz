        package main
        import "fmt"

        func mostrar_vetor(lista []int, sep string) {
             fmt.Print("[")
            for i, valor := range lista {
                if i != 0 {
                    fmt.Print(", ")
                }
                fmt.Printf("%v ", valor)
            }
            fmt.Print("]\n")
        }
        func main() { //lista estática
             var qtd int
             fmt.Scan(&qtd)
             lista := []int {1, 2, 3}
            var idades []int = make([]int, qtd)
            for i := range lista {

            }
            fmt.Println(idades)
            lista := []int{9, 8, 7, 6, 5, 2, 1, 1}
           mostrar_vetor(lista, ", ")
        }