/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 21:14
* Developer : ibrahimcobani
*
*******/
package multipleReturnValue

import "fmt"

type Item struct {
	no            string
	name          string
	unitofmeasure string
	unitprice     float32
}

func GetItemInfo(paramItem Item) {
	firstReturnValue, lastReturnValue := CreateMessage(paramItem.no, paramItem.name)
	fmt.Println(firstReturnValue, lastReturnValue)
}

func CreateMessage(code, name string) (string, string) {
	return "{" + "code:'" + code + "', name: '" + name + "'}", " --- *** --- Bu da ikinci parametre " + code
}

func main() {
	var item1 = Item{"LAN001", "Pencil", "Adet", 3}
	GetItemInfo(item1)
}
