/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 21:04
* Developer : ibrahimcobani
*
*******/
package addingaReturn

import "fmt"

type Item struct {
	no            string
	name          string
	unitofmeasure string
	unitprice     float32
}

func GetItemInfo(paramItem Item) {
	fmt.Println(CreateMessage(paramItem.no, paramItem.name))
}

func CreateMessage(code, name string) string {
	return "{" + "code:'" + code + "', name: '" + name + "'}"
}

func main() {
	var item1 = Item{"LAN001", "Pencil", "Adet", 3}
	GetItemInfo(item1)
}