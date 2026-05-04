            package main
            import "fmt"
            func main() {
                var prof, salto, escorr, i int
                fmt.Scan(&prof)
                fmt.Scan(&salto)
                fmt.Scan(&escorr)

                for i := 0; i < prof; i+= salto {
                    if i >= prof {
                        fmt.Println("saiu")
                    } else {
                        i -= escorr
                    }

                }
                fmt.Println(i, i)
        }