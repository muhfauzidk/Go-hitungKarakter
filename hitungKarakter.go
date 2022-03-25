package main

import "fmt"

func main() {
	var x string
	var y int

	y = 1

	fmt.Print("Masukkan kata: ", x)
	fmt.Scanln(&x)

	//Menghitung jumlah huruf yang diinputkan
	fmt.Println("Jumlah karakter : ", len(x))

	//Mencari karakter pada posisi tertentu
	fmt.Print("Posisi karakter : ")
	fmt.Scan(&y)
	fmt.Println("adalah : ", string(x[y]))
}
