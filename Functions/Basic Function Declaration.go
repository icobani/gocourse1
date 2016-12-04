/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 20:32
* Developer : ibrahimcobani
*
*******/
package basicFunctionDecleration

import "fmt"

type Item struct {
	no            string
	name          string
	unitofmeasure string
	unitprice     float32
}

func GetItemPrice(paramItem Item) {
	fmt.Println(paramItem.unitprice)
}

func main() {
	var item1 = Item{"LAN001", "Pencil", "Adet", 3}
	GetItemPrice(item1)
}
