package main

import "fmt"

func mainscanner () {
	var nama string 
	fmt.Println("Masukan Nama Anda : ")
	fmt.Scanf("%s", &nama)
	fmt.Printf("Nama yang di input adalah : %s \n",nama)


	var uang uint
	fmt.Println("Masukan Jumlah Uang yang ingin anda Kirim : ")
	fmt.Scanf("%d", &uang)
	fmt.Printf("Uang yang masuk adalah : %d \n",uang)

	var suhu float32

	fmt.Println("Masukan Jumlah Suhu Sekarang : " )
	fmt.Scanf("%f",&suhu)
	fmt.Printf("Suhu Sekarang Adalah : %f \n",suhu)






}