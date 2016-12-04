/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 01/12/2016    
* Time      : 19:49
* Developer : ibrahimcobani
*
*******/
package Closures

import "fmt"

type Printer func(string) ()

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(custom + s)
	}
}

func Great(name string, do Printer) {
	do(name)
}

/*func main() {
	Great("ibrahim", CreatePrintFunction("Merhaba "))
}

*/