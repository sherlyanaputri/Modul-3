package main
import "fmt"

func main() {
	var alas, tinggi int

	fmt.Scan(&alas, &tinggi)

	luas := (alas * tinggi) / 2

	fmt.Println(luas)
}