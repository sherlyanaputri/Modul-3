package main
import "fmt"

func main() {
    var gajiPokok, lembur int
    var bonus, potongan, gajiBersih int

    fmt.Scan(&gajiPokok, &lembur)

    bonus = 45000 * lembur
    potongan = gajiPokok * 55 / 1000
    gajiBersih = gajiPokok + bonus - potongan

    fmt.Println(gajiBersih)
}