package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Print("Masukkan nama, nim, kelas: ")
	fmt.Scanln(&nama, &nim, &kelas)

	fmt.Println("Perkenalkan saya adalah " + nama + ", salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + " dengan NIM " + nim + ".")
}