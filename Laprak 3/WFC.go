package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    hari := (4+n-1)%7 + 1

    fmt.Println(hari)
}