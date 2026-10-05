package main

import "fmt"

func main() {
	var f, c float64

	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scanln(&f)

	c = (f - 32) * 5 / 9

	fmt.Println("Suhu Celcius =", c)
}