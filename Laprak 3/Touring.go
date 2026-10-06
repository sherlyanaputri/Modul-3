package main
import "fmt"

func main() {

    // [DEKLARASI VARIABEL]
    var kecepatan, menit int
    var jam, total_jarak int

    // [INPUT]
    fmt.Scan(&kecepatan)

    // [PROSES]
    total_jarak = 100 + 60 + 170
    menit = (total_jarak * 60) / kecepatan
    jam = menit / 60
    menit = menit % 60

    // [OUTPUT]
    fmt.Println(jam, menit)
}