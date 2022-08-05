package main

import "fmt"

func konversinilai(nilai int) string {
	var hasil string

	switch {
	case nilai >= 80 && nilai <= 100:
		hasil = "A"
	case nilai >= 65 && nilai <= 80:
		hasil = "B+"
	case nilai >= 50 && nilai <= 64:
		hasil = "B"
	case nilai >= 35 && nilai <= 49:
		hasil = "C"
	case nilai >= 0 && nilai <= 34:
		hasil = "D"
	default:
		hasil = "Nilai tidak boleh lebih atau kurang dari nilai 0-100"
	}
	return hasil
}
func main() {
	var nilai int = 101
	fmt.Println(konversinilai(nilai))
}
