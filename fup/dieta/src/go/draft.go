package main
import "fmt"
func main() {
    var dias int
    fmt.Scan(&dias)

    cal := make([]int, dias) //criar uma lists que armazene os dias
    for i := 0; i < dias ; i++ { //para i = 0, enquanto i for menor que dias, i incrementa
        fmt.Scan(&cal[i]) //escanear os índices da lista
    }
    soma := 0
    for _, i2 := range cal { 
        soma += i2
    }
    media := soma / dias
    fmt.Printf("%.1f\n", float64(media))

    }

