package main
import "fmt"
func main() {
    var tempo , horas, minutos, segundos int
    
    fmt.Scan(&tempo) //ler o tempo
    fmt.Scan(&horas) //ler as horas
    fmt.Scan(&minutos) //ler os minutos
    fmt.Scan(&segundos) //ler os segundos
    
    horas = (tempo/3600) //conseguir as horas
    minutos = (tempo%3600)/ 60 //conseguir os minutos
    segundos = (tempo%3600)%60 //conseguir os segundos
    fmt.Printf("%d:%d:%d\n", horas, minutos, segundos) //imprimir a formatação de horas, minutos e segundos
}
