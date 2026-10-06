package main
import "fmt"

func main() {
    var koin, emas, perak, tembaga int

    fmt.Scan(&koin)

    emas = koin / 9
    sisa := koin % 9
    perak = sisa / 3
    tembaga = sisa % 3

    fmt.Println(emas, perak, tembaga)
}