/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 17:58
* Developer : ibrahimcobani
*
*******/
package main

type Item struct {
	no            string
	name          string
	unitofmeasure string
	unitprice     float32
}

func main() {
	var item1 = Item{"LAN001", "Pencil", "Adet", 3}
	println(item1.no, item1.name, item1.unitofmeasure, item1.unitprice)

	item1 = Item{no:"LAN003", name:"Bike", unitofmeasure:"ADET", unitprice:4.56}
	println(item1.no, item1.name, item1.unitofmeasure, item1.unitprice)

	item1.no = "LAN004"
	item1.name = "car"
	item1.unitofmeasure = "KOLI"
	item1.unitprice = 4.564
	println(item1.no, item1.name, item1.unitofmeasure, item1.unitprice)
}
