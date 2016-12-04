/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 29/11/2016    
* Time      : 22:22
* Developer : ibrahimcobani
*
*******/
package NamedReturns

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

func CreateMessage(code, name string) (Code string, Name string) {
	Code = "{" + "code:'" + code + "', name: '" + name + "'}"
	Name = " --- *** --- Bu da ikinci parametre " + code
	return
}

func main() {
	var item1 = Item{"LAN001", "Pencil", "Adet", 3}
	GetItemInfo(item1)
}
