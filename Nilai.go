package main

import "fmt"

func rataRata(number...int) float64 {
	total := 0
	for _, value := range number {
		total += value
	}
	return float64(total) / float64(len(number))
}

func nilaiTertinggi(number... int) int {
	nilai := number[0]
	for _, value := range number {
		if value > nilai {
			nilai = value
		}
	}
	return nilai
}

func jumlah(number...int) int {
	return len(number)
}

func main () {
	arr := [] int {} 
	var angka, inputan int
	
	fmt.Print("Masukkan Jumlah Inputan: ")
	fmt.Scanln(&inputan)

	for i := 1; i <= inputan; i++ {	
		fmt.Printf("Masukkan Angka ke-%d: ", i)
		fmt.Scanln(&angka)
		arr = append(arr, angka)
	}  


	
	fmt.Println("Nilai Rata-Rata Adalah: ", rataRata(arr...))
	fmt.Println("Nilai Tertinggi Adalah: ", nilaiTertinggi(arr...))
	fmt.Println("Jumlah Angka Yang Dimasukkan: ", jumlah(arr...) )
}