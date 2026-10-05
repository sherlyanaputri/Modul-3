package main

import "fmt"

func main() {
	var ngueng int = 15
	alamatmemori := &ngueng

	fmt.Println(alamatmemori)
	fmt.Println(*alamatmemori)
}