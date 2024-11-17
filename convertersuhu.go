package main

import "fmt"


type hitungSuhu interface {
	toCelcius() float64
	toKelvin() float64
	toFahrenheit() float64
}
type Fahrenheit struct {
	suhu float64
}
type Celcius struct {
	suhu float64
} 
type Kelvin struct {
	suhu float64
}

func (c Celcius) toCelcius() float64{
	return c.suhu
}
func (c Celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32 
}
func (c Celcius) toKelvin() float64 {
	return c.suhu + 273.15
}
func (k Kelvin) toKelvin() float64 {
	return k.suhu
}
func (k Kelvin) toCelcius() float64{
	  return k.suhu - 273.15
}
func (k Kelvin) toFahrenheit() float64{
	return  (k.suhu * (9.0 / 5.0 )) - 459.67 
}
func (f Fahrenheit) toFahrenheit() float64 {
	return f.suhu
}
func (f Fahrenheit) toCelcius() float64 {
	return (f.suhu - 32 ) *(5.0 / 9.0)
}
func (f Fahrenheit) toKelvin () float64 {
	return  (f.suhu + 459.67) * (5.0 / 9.0)
}


func main () {
 fmt.Println("Pilih Suhu Awal :")
 fmt.Println("1.Celcius")
 fmt.Println("2.Kelvin")
 fmt.Println("3.Fahrenheit")
 fmt.Println("Masukan Suhu Awal yang di inginkan :  ")

 var suhuAwal int 
 fmt.Scanf("%d", &suhuAwal)
 for suhuAwal  < 1 || suhuAwal > 3 {
	fmt.Println(" Ga valid cuy , Masukan suhu yang di inginkan : ")
	fmt.Scanf("%d",&suhuAwal)
}
fmt.Println("Pilihan suhu akhir : ")
fmt.Println("1.Celcius")
fmt.Println("2.Kelvin")
fmt.Println("3.Fahrenheit")
fmt.Println("Masukan Suhu akhir yang di inginkan : ")

var suhuAkhir int 
fmt.Scanf("%d",&suhuAkhir)
 for suhuAkhir  < 1 || suhuAkhir > 3 {
	fmt.Println(" Ga valid cuy , Masukan suhu yang di inginkan : ")
	fmt.Scanf("%d",&suhuAkhir)
}

var suhu float64 
fmt.Println("Masukan Suhu nya Cuy :  ")
fmt.Scanf("%f",&suhu)

var interfaceSuhu hitungSuhu

switch suhuAwal {
case 1 :
	interfaceSuhu = Celcius{suhu}
case 2 :
	interfaceSuhu = Kelvin{suhu}
case 3 : 
	interfaceSuhu = Fahrenheit{suhu}
} 

var suhuFinal float64 
switch suhuAkhir {
case 1 : 
	suhuFinal = interfaceSuhu.toCelcius()
case 2 : 
suhuFinal = interfaceSuhu.toKelvin()
case 3 : 
	 suhuFinal = interfaceSuhu.toFahrenheit()
}
fmt.Printf("suhu akhir yang di dapatkan adalah : %.2f",suhuFinal)






}
