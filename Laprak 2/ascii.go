package main

import "fmt"

func main() {
    // Baris 1: baca 5 integer, cetak sebagai karakter
    var a [5]int
    for i := 0; i < 5; i++ {
        fmt.Scan(&a[i])
    }
    for i := 0; i < 5; i++ {
        fmt.Printf("%c", a[i])
    }
    fmt.Println()

    // Baris 2: baca 3 karakter berdampingan, cetak karakter setelahnya
    var s string
    fmt.Scan(&s)
    for _, c := range s {
        fmt.Printf("%c", c+1)
    }
    fmt.Println()
}