package main

import "fmt"

func main() {
	var nilai = 80
	var absensi = 80

	var lulusUjian = nilai >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilai >= 80 && absensi >= 80)
}